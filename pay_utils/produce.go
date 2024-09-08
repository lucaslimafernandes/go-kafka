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

	for {
		inputer()
	}
}

func inputer() {

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

		go func(genSell models.Sell) {
			defer wg.Done()
			getResponse(genSell)
		}(genSell)

	}

	wg.Wait()

}

func getResponse(s models.Sell) {

	for {
		msg, err := models.CMR.ReadMessage(-1)
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			continue
		}

		// Deserializar a mensagem em uma struct 'response'
		var resp models.ResponseSell
		err = json.Unmarshal(msg.Value, &resp)
		if err != nil {
			log.Printf("Failed to unmarshal message: %v\n", err)
			continue
		}

		// Filtrar as mensagens com base nos campos desejados
		if shouldProcess(s, resp) {
			log.Printf("Mensagem processada: %+v\n", resp)
		}

	}

}

// Função para aplicar a lógica de filtragem
func shouldProcess(s models.Sell, r models.ResponseSell) bool {
	// Exemplo de filtro básico: verificar se PersonId > 0 e Amount > 100
	// return resp.PersonId > 0 && resp.Amount > 100.00 && strings.Contains(resp.Address.City, "New York")
	return r.PersonId == s.PersonId && r.Amount == s.Amount && r.Address == s.Address
}
