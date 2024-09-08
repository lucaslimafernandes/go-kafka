package controllers

import (
	"encoding/json"
	"log"
	"math/rand"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/lucaslimafernandes/go-kafka-prd/models"
)

func PassingCards() {

	for {
		inputData()
	}

}

func inputData() {

	var wg sync.WaitGroup
	topic := "Sells"
	sells := rand.Intn(1000)

	for i := 0; i <= sells; i++ {

		wg.Add(2)

		genSell, err := models.Selling()
		if err != nil {
			return
		}
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
