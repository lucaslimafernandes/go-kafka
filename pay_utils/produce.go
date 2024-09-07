package payutils

import (
	"encoding/json"
	"log"
	"math/rand"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/lucaslimafernandes/go-kafka/models"
)

func PassingCards() {

	for i := 0; i <= 10; i++ {
		inputer()
	}
}

func inputer() {

	var wg sync.WaitGroup
	topic := "Sells"
	sells := rand.Intn(100)
	for i := 0; i <= sells; i++ {

		wg.Add(1)

		go func() {

			genSell, err := models.Selling()
			if err != nil {
				return
			}

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

		}()

		wg.Wait()

	}
}
