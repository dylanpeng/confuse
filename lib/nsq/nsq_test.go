package nsq

import (
	"confuse/lib/logger"
	"github.com/nsqio/go-nsq"
	"strconv"
	"testing"
	"time"
)

var Log logger.ILogger

type MessageHandle struct {
}

func (h *MessageHandle) HandleMessage(message *nsq.Message) error {
	Log.Infof("receive message. | id: %d | body: %s | addr: %s", message.ID, string(message.Body), message.NSQDAddress)
	return nil
}

func TestNsqProducerConsumer(t *testing.T) {
	var err error
	Log, err = logger.NewLogger(&logger.Config{
		FilePath:   "./logs/confuse",
		Level:      "debug",
		TimeFormat: "2006-01-02 15:04:05.000",
		MaxAgeDay:  30,
	})

	if err != nil {
		t.Fatalf("NewLogger fail. | err: %s", err)
	}

	consumerConf := &ConsumerConfig{
		Topic:      "nsq-test",
		Channel:    "ch-nsq-test",
		LookupAddr: []string{"127.0.0.1:4161"},
	}

	consumer, err := NewConsumer(consumerConf, Log, &MessageHandle{})
	if err != nil {
		t.Fatalf("NewConsumer fail. | err: %s", err)
	}

	producerConfig := &ProducerConfig{
		NsqdAddr: []string{"127.0.0.1:4150"},
	}

	producer, err := NewProducer(producerConfig, Log)
	if err != nil {
		t.Fatalf("NewProducer fail. | err: %s", err)
	}

	for i := 0; i < 100; i++ {
		producer.SendMessage("nsq-test", strconv.Itoa(i))
		//time.Sleep(100 * time.Millisecond)
	}

	time.Sleep(1 * time.Minute)
	consumer.Stop()
	producer.Stop()
}
