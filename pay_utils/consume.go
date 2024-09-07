package payutils

import (
	"context"
	"encoding/json"
	"log"

	"github.com/lucaslimafernandes/go-kafka/models"
)

func Validat(msg string) {

	var sell models.Sell
	var sellLog models.SellLog
	userAcc := models.GetUser(1)

	json.Unmarshal([]byte(msg), &sell)

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

	insertLog(sellLog)

}

func insertLog(s models.SellLog) {

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
		log.Printf("Failed to insert new user: %v\n", err)
	}
}

func haveBalance() {

}
