package payutils

import (
	"context"
	"log"
	"math/rand"
	"sync"

	"github.com/lucaslimafernandes/go-kafka/models"
)

var wg sync.WaitGroup

func Inserts(n, g int) {

	for t := 0; t <= g; t++ {
		onePerson(n)
	}

}

func onePerson(n int) {

	for t := 0; t <= n; t++ {

		wg.Add(1)

		go func() {

			user, _ := models.GeneratorData()
			min := 1000.00
			max := 20000.00
			balance := min + rand.Float64()*(max-min)

			insertQuery := `
				INSERT INTO users (name, credit_card, balance, address, city, state, postal_code, coordinates_lat, coordinates_long)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, &9); 
			`
			err := models.DB.QueryRow(context.Background(),
				insertQuery,
				user.Name,
				user.CreditCard,
				balance,
				user.Address.Address,
				user.Address.City,
				user.Address.State,
				user.Address.PostalCode,
				user.Address.Coordinates.Latitude,
				user.Address.Coordinates.Longitude,
			).Scan()
			if err != nil && err.Error() != "no rows in result set" {
				log.Printf("Failed to insert new user: %v\n", err)
			}

			wg.Done()

		}()

		wg.Wait()

	}
}
