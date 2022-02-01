package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"fmt"
)

func main() {
	configMap := kafka.ConfigMap{
		"bootstrap.servers":"kafka:9092",
		"group.id":"mygroup",
	}
	consumer, err := kafka.NewConsumer(&configMap)
	if err != nil {
		log.Fatal(err)
	}
	consumer.Subscribe("exemplo",nil)
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value))
		}
	}
}