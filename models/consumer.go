package models

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var CM *kafka.Consumer

func ConnectKafkaConsumer() {

	var err error

	CM, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "my-consumer-group", // Identificador do grupo de consumidores
		"auto.offset.reset": "earliest",          // Começa a ler a partir do primeiro offset disponível
	})

	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %s\n", err)
	}

	err = CM.Subscribe("Sells", nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s\n", err)
	}

	// Loop para consumir as mensagens
	fmt.Println("Waiting for messages...")

}
