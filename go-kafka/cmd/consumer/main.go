package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "fc2-gokafka-kafka-1:9092",
		"client.id":         "goapp4-consumer",
		"group.id":          "goapp-group",
		"auto.offset.reset": "earliest", // earliest = Consumidor le todas as mensagens desde o inicio.
	}

	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("error consumer", err.Error())
	}

	topics := []string{"teste"}
	c.SubscribeTopics(topics, nil)

	for {
		msg, err := c.ReadMessage(-1) //-1 = para timeout ficar conectado para sempre
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		}
	}
}
