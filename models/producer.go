package models

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var PD *kafka.Producer

func ConnectKafkaProducer() {

	var err error
	PD, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		log.Fatalln("Failed to connect kafka: ", err)
	}

	// monitoring

	go func() {

		for e := range PD.Events() {

			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Delivery failed: %s\n", ev.TopicPartition)
				} else {
					log.Printf("Delivered message to: %s\n", ev.TopicPartition)
				}
			}
		}
	}()

}
