package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChan := make(chan kafka.Event)
	producer := NewKafkaProducer()

	Publish("transferiu1", "teste", producer, []byte("transferecia2"), deliveryChan)
	DeliveryReport(deliveryChan) // sync

	//go DeliveryReport(deliveryChan) // async
	//producer.Flush(5000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "fc2-gokafka-kafka-1:9092",
		"delivery.timeout.ms": "0",
		"acks":                "all",  //[0 = Manda a mensgem e não espera retorno, 1 = Aguardo retorno pelo menos se o lider persistiu a mensagem, all = Espero o lider e todos os brokers tenham persistido a mensagem]
		"enable.idempotence":  "true", //Garantir que a mensagem é entregue na ordem que vc quer e somente uma vez, em troca de performance. Para idempontence estar ativa "acks" deve ser "all".
	}
	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar")
			} else {
				fmt.Println("Mensagem enviada:", ev.TopicPartition)
				// Possivel caso de uso:
				// anotar no banco de dados que a mensagem foi processado.
				// ex: confirma que uma transferencia bancaria ocorreu.
			}
		}
	}
}
