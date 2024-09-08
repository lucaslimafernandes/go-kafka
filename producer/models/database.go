package models

import (
	"context"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool
var PD *kafka.Producer

func ConnectDB() {

	var err error
	var_db := "user=postgres password=password host=localhost port=5432 dbname=postgres"

	DB, err = pgxpool.New(context.Background(), var_db)
	if err != nil {
		log.Fatalln("Failed to connect to the database: ", err)
	}

}

func ConnectKafkaConsumer() {

	var err error

	PD, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		log.Fatalln("Failed to connect kafka: ", err)
	}

	// monitoring events
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
