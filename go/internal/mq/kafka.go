package mq

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/beckbikang/WriteCodeEveryDay/global"
	"github.com/segmentio/kafka-go"
)

var (
	KafkaWriter *kafka.Writer
	KafkaReader *kafka.Reader
)

func SetupKafkaProducer() error {
	KafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP(global.KafkaSetting.Brokers...),
		Topic:    global.KafkaSetting.Topic,
		Balancer: &kafka.LeastBytes{},
	}
	
	return nil
}

func SetupKafkaConsumer() error {
	KafkaReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers: global.KafkaSetting.Brokers,
		Topic:   global.KafkaSetting.Topic,
		GroupID: global.KafkaSetting.Group,
	})
	
	return nil
}

func CloseKafkaProducer() {
	if KafkaWriter != nil {
		KafkaWriter.Close()
	}
}

func CloseKafkaConsumer() {
	if KafkaReader != nil {
		KafkaReader.Close()
	}
}

func SendMessage(message interface{}) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	
	return KafkaWriter.WriteMessages(context.Background(), kafka.Message{
		Value: data,
	})
}

func ConsumeMessages(handler func(message []byte) error) {
	for {
		msg, err := KafkaReader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Kafka consume error: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}
		
		err = handler(msg.Value)
		if err != nil {
			log.Printf("Message handler error: %v", err)
		}
	}
}
