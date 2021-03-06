package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	producer := NewKafkaProducer()
	Publish("mensagem", "teste", producer, nil)
	producer.Flush(1000)
}

func NewKafkaProducer() *kafka.Producer{
	configmap := &kafka.ConfigMap{
		"bootstrap.servers": "gokafka_kafka_1:9092",
	}
	p, err := kafka.NewProducer(configmap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

func Publish(msg, topic string, producer *kafka.Producer, key []byte) error {
	message := &kafka.Message{
		Value: []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key: key,
	}
	err := producer.Produce(message, nil)
	if err != nil {
		return err
	}
	return nil
}