package main

import "github.com/lucaslimafernandes/go-kafka-prd/models"

// import (
// 	"flag"
// 	"fmt"

// )

func init() {

	models.ConnectDB()

	models.ConnectKafkaConsumer()
	// models.ConnectKafkaConsumerResp()

}

// func main() {

// 	newUsersFlag := flag.Bool("new_users", false, "new users")
// 	runFlag := flag.Bool("run", false, "run")

// 	flag.Parse()

// 	if *newUsersFlag {
// 		payutils.Inserts(1000, 1000)
// 	}

// 	if *runFlag {
// 		fmt.Println("Vai executar")
// 		go payutils.PassingCards()
// 	}

// 	fmt.Println("Terminou")
// 	fmt.Println("Consuming!")

// 	consumer()

// }

// func consumer() {

// 	for {
// 		msg, err := models.CM.ReadMessage(-1) // Lê mensagem (timeout infinito)
// 		if err == nil {
// 			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
// 			payutils.Validat(string(msg.Value))
// 		} else {
// 			// Acontece se houver um erro ao ler a mensagem
// 			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
// 		}
// 	}

// }
