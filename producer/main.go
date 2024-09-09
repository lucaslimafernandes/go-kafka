package main

import (
	"flag"
	"fmt"

	"github.com/lucaslimafernandes/go-kafka-prd/controllers"
	"github.com/lucaslimafernandes/go-kafka-prd/models"
)

// import (
// 	"flag"
// 	"fmt"

// )

func init() {

	models.ConnectMongoDB()
	models.ConnectKafkaProducer()

}

func main() {

	newUsersFlag := flag.Bool("new_users", false, "new users")
	runFlag := flag.Bool("run", false, "run")
	timeRunning := flag.Duration("timer", 1, "Time running")

	flag.Parse()

	if *newUsersFlag {
		models.GenNewUsers()
	}

	if *runFlag {
		fmt.Println("Vai executar")
		controllers.PassingCards(*timeRunning)
	}

	// 	fmt.Println("Terminou")
	// 	fmt.Println("Consuming!")

	// 	consumer()

}

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
