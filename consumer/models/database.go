package models

import (
	"context"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var CM *kafka.Consumer
var Mongo *mongo.Client

func ConnectKafkaConsumer() {

	var err error

	CM, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "my-consumer-group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %s\n", err)
	}

	err = CM.Subscribe("Sells", nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s\n", err)
	}

	fmt.Println("Waiting for messages...")

}

func ConnectMongoDB() {

	var err error

	uri := "mongodb://mongoadmin:secret@localhost:27017/"
	clientOptions := options.Client().ApplyURI(uri)

	Mongo, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalln("Failed to connect MongoDB: ", err)
	}

	err = Mongo.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("MongoDB connection established successfully!")

}
