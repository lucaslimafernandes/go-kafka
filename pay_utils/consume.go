package payutils

import (
	"context"
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/lucaslimafernandes/go-kafka/models"
)

func Validat(msg string) {

	var sell models.Sell
	var sellLog models.SellLog
	json.Unmarshal([]byte(msg), &sell)

	userAcc := models.GetUser(sell.PersonId)

	if userAcc.Balance >= sell.Amount {
		sellLog = models.SellLog{
			User_id: sell.PersonId,
			Amount:  sell.Amount,
			City:    sell.Address.City,
			State:   sell.Address.State,
			IsValid: true,
		}
	} else {
		sellLog = models.SellLog{
			User_id: sell.PersonId,
			Amount:  sell.Amount,
			City:    sell.Address.City,
			State:   sell.Address.State,
			IsValid: false,
		}
	}

	insertLog(sellLog, userAcc.Balance-sell.Amount)
	response(sell, sellLog.IsValid)

}

func insertLog(s models.SellLog, new_balance float64) {

	insertQuery := `
		INSERT INTO logs (user_id, amount, city, state, is_valid)
		VALUES ($1, $2, $3, $4, $5);
	`

	err := models.DB.QueryRow(context.Background(),
		insertQuery,
		s.User_id,
		s.Amount,
		s.City,
		s.State,
		s.IsValid,
	).Scan()
	if err != nil && err.Error() != "no rows in result set" {
		log.Printf("Failed to insert new purchase: %v\n", err)
	}

	if new_balance >= 0.0 {
		updateQuery := `
			UPDATE users 
			SET balance = $1
			WHERE id = $2;
		`

		_, err = models.DB.Exec(context.Background(),
			updateQuery,
			new_balance,
			s.User_id,
		)
		if err != nil {
			log.Fatalf("Failed to update balance to user: %s\n", err)
		}
	}

}

func response(s models.Sell, valid bool) {

	res := models.ResponseSell{
		PersonId: s.PersonId,
		Amount:   s.Amount,
		Address:  s.Address,
		IsValid:  valid,
	}

	topic := "Response"
	jsonRes, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("Failed response: %v\n", err)
	}

	err = models.PD.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topic, Partition: kafka.PartitionAny,
		},
		Value: jsonRes,
	}, nil)
	if err != nil {
		log.Printf("Failed to produce message: %v\n", err)
		return
	}

}
