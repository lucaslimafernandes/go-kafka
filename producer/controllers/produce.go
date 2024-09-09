package controllers

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/lucaslimafernandes/go-kafka-prd/models"
)

func PassingCards(t time.Duration) {

	timeout := time.After(t * time.Minute)

	for {
		select {
		case <-timeout:
			log.Println("Timeout reached. Terminating execution.")
			os.Exit(0)
		default:
			inputData()
		}
	}

}

func inputData() {

	var wg sync.WaitGroup
	topic := "Sells"
	sells := rand.Intn(1000)

	for i := 0; i <= sells; i++ {

		wg.Add(1)

		genSell := models.Selling()
		go func(genSell models.Sell) {

			jsonSell, err := json.Marshal(genSell)
			if err != nil {
				log.Printf("Failed to marshal data: %s\n", err)
				return
			}

			err = models.PD.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic: &topic, Partition: kafka.PartitionAny,
				},
				Value: jsonSell,
			}, nil)
			if err != nil {
				log.Printf("Failed to produce message: %v\n", err)
				return
			}

			wg.Done()

		}(genSell)

		// go func(genSell models.Sell) {
		// 	defer wg.Done()
		// 	getResponse(genSell)
		// }(genSell)

	}

	wg.Wait()

}
