package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lucaslimafernandes/go-kafka-cons/controllers"
	"github.com/lucaslimafernandes/go-kafka-cons/models"
)

func init() {

	models.ConnectMongoDB()
	models.ConnectKafkaConsumer()
}

func main() {

	timeExec := flag.Duration("timer", 1, "Time of execution (min)")

	flag.Args()

	consumer(*timeExec)
}

func consumer(d time.Duration) {

	timeout := time.After(d * time.Minute)

	for {

		select {
		case <-timeout:
			log.Println("Timeout reached. Terminating execution.")
			os.Exit(0)
		default:
			msg, err := models.CM.ReadMessage(-1) // Lê mensagem (timeout infinito)
			if err == nil {
				fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
				controllers.Validate(string(msg.Value))
			} else {
				// Acontece se houver um erro ao ler a mensagem
				fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			}
		}
	}

}
