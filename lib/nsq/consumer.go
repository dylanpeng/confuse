package nsq

import (
	"confuse/lib/logger"
	"github.com/nsqio/go-nsq"
)

type ConsumerConfig struct {
	Topic      string   `toml:"topic" json:"topic"`
	Channel    string   `toml:"channel" json:"channel"`
	LookupAddr []string `toml:"lookup_addr" json:"lookup_addr"`
}

type Consumer struct {
	c       *ConsumerConfig
	logger  logger.ILogger
	client  *nsq.Consumer
	handler nsq.Handler
}

func NewConsumer(c *ConsumerConfig, logger logger.ILogger, handle nsq.Handler) (*Consumer, error) {
	consumer := &Consumer{
		c:       c,
		logger:  logger,
		handler: handle,
	}

	// 1. 创建消费者
	config := nsq.NewConfig()
	client, err := nsq.NewConsumer(c.Topic, c.Channel, config)
	if err != nil {
		logger.Errorf("init nsq consumer failed. | config: %+v | err: %s", c, err)
		return nil, err
	}

	// 2. 添加处理消息的方法
	client.AddHandler(handle)
	consumer.client = client

	go consumer.run()

	return consumer, nil
}

func (c *Consumer) run() {
	// 3. 通过http请求来发现nsqd生产者和配置的topic（推荐使用这种方式）
	err := c.client.ConnectToNSQLookupds(c.c.LookupAddr)
	if err != nil {
		c.logger.Errorf("ConnectToNSQLookupds failed. | config: %+v | err: %s", c.c, err)
		return
	}

	// 4. 接收消费者停止通知
	<-c.client.StopChan

	// 5. 获取统计结果
	stats := c.client.Stats()
	c.logger.Infof("message received %d, finished %d, requeued:%s, connections:%s",
		stats.MessagesReceived, stats.MessagesFinished, stats.MessagesRequeued, stats.Connections)
}

func (c *Consumer) Stop() {
	c.client.StopChan <- 0
	c.logger.Infof("nsq consumer stop. | config: %+v", c.c)
}
