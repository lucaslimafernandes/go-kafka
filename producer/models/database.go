package models

import (
	"context"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *pgxpool.Pool
var PD *kafka.Producer
var Mongo *mongo.Client

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
