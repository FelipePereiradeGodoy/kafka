package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

//Crie um novo consumer "isolado", com outro group-id (para não interferir no consumer real)
//Após fazer isso substitua abaixo o nome do topico, nome do group-id e range das mensagens e a particação especifica.

func main() {
	topic := "topic_name"
	partition := 0 // especfic partition
	startOffset := int64(4)
	endOffset := int64(8)

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "group-id-to-retry",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatal(err)
	}

	defer consumer.Close()

	err = consumer.Assign([]kafka.TopicPartition{
		{
			Topic:     &topic,
			Partition: int32(partition),
			Offset:    kafka.Offset(startOffset),
		},
	})

	if err != nil {
		log.Fatal("Erro ao fazer Assign:", err)
	}

	fmt.Println("Reprocessando offsets de", startOffset, "até", endOffset)

	for {
		msg, err := consumer.ReadMessage(2 * time.Second)

		if err != nil {
			fmt.Println("Timeout esperando mensagens ...")
			continue
		}

		offset, _ := strconv.ParseInt(msg.TopicPartition.Offset.String(), 10, 64)

		if offset > endOffset {
			fmt.Println("Fim do range alcançado. Encerrando ...")
			break
		}

		fmt.Printf("Reprocessando offset %d: %s\n", offset, string(msg.Value))
	}

	fmt.Println("Reprocessamento concluído.")
}
