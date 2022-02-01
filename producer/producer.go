package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":"kafka:9092",
	}

	producer, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Fatal("Erro ao configurar producer")
	}

	topic := "exemplo"
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value: []byte("Veio do Go!"),
		Key: []byte("mykey"),
	}
	err = producer.Produce(msg, nil)
	if err != nil {
		log.Fatal("Erro ao produzir")
	}
	forever := make(chan bool)
	<- forever
}