package main

import (
	"flag"
	"fmt"

	"github.com/lucaslimafernandes/go-kafka-prd/controllers"
	"github.com/lucaslimafernandes/go-kafka-prd/models"
)

func init() {

	models.ConnectMongoDB()
	models.ConnectKafkaProducer()

}

func main() {

	newUsersFlag := flag.Bool("new_users", false, "new users")
	runFlag := flag.Bool("run", false, "run")
	timeRunning := flag.Duration("timer", 1, "Time running (min)")

	flag.Parse()

	if *newUsersFlag {
		models.GenNewUsers()
	}

	if *runFlag {
		fmt.Println("Vai executar")
		controllers.PassingCards(*timeRunning)
	}

}
