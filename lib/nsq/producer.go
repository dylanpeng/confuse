package nsq

import (
	"confuse/lib/logger"
	"encoding/json"
	"errors"
	"github.com/nsqio/go-nsq"
)

type ProducerConfig struct {
	NsqdAddr []string `toml:"nsqd_addr" json:"nsqd_addr"`
}

type Producer struct {
	c         *ProducerConfig
	logger    logger.ILogger
	producers []*nsq.Producer
}

func NewProducer(c *ProducerConfig, logger logger.ILogger) (*Producer, error) {
	producer := &Producer{
		c:         c,
		logger:    logger,
		producers: make([]*nsq.Producer, 0, 8),
	}

	config := nsq.NewConfig() // 1. 创建生产者
	nsqProducer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		logger.Errorf("NewProducer fail. | config: %+v | err: %s", config, err)
		return nil, err
	}

	err = nsqProducer.Ping() // 2. 生产者ping
	if err != nil {
		logger.Errorf("NewProducer Ping fail. | config: %+v | err: %s", config, err)
		return nil, err
	}

	producer.producers = append(producer.producers, nsqProducer)

	return producer, nil
}

func (p *Producer) SendMessage(topic string, message interface{}) error {
	msg, err := json.Marshal(message)
	if err != nil {
		p.logger.Warningf("SendMessage fail. | topic: %s | message: %+v | err: %s", topic, message, err)
		return err
	}

	if len(p.producers) == 0 {
		err = errors.New("producer not init")
		p.logger.Warningf("producer not init fail. | topic: %s | message: %+v | err: %s", topic, message, err)
		return err
	}

	err = p.producers[0].Publish(topic, msg) // 注意one-test　对应消费者consumer.go　保持一致
	if err != nil {
		p.logger.Warningf("Publish fail. | topic: %s | message: %+v | err: %s", topic, message, err)
		return err
	}

	return nil
}

func (p *Producer) Stop() {
	for _, nsqProducer := range p.producers {
		nsqProducer.Stop()
	}
}
