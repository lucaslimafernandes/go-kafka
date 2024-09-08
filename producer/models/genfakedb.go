package models

import (
	"context"
	"log"
	"math/rand"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
)

var wg sync.WaitGroup

func GenNewUsers() {

	for t := 0; t <= 1000; t++ {
		onePerson()
	}

}

func onePerson() {

	collection := Mongo.Database("public").Collection("users")

	for i := 0; i <= 1000; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()

			user, _ := GeneratorData()
			min := 1000.00
			max := 20000.00
			balance := min + rand.Float64()*(max-min)

			document := bson.M{
				"name":        user.Name,
				"credit_card": user.CreditCard,
				"balance":     balance,
				"address":     user.Address.Address,
				"city":        user.Address.City,
				"state":       user.Address.State,
				"postal_code": user.Address.PostalCode,
				"coordinates": bson.M{
					"lat":  user.Address.Coordinates.Latitude,
					"long": user.Address.Coordinates.Longitude,
				},
			}

			_, err := collection.InsertOne(context.Background(), document)
			if err != nil {
				log.Printf("Failed to insert new user: %s\n", err)
			}

		}()
	}
	wg.Wait()

	log.Println("New users are inserted!")
}
