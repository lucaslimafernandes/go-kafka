package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/lucaslimafernandes/go-kafka-cons/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Validate(msg string) {

	var sell models.Sell
	var sellLog models.SellLog
	json.Unmarshal([]byte(msg), &sell)

	fmt.Printf("\n\nUser Balance: %f, Sell Amount: %f\n\n", sell.Person.Balance, sell.Amount)

	if sell.Person.Balance > 0 {
		fmt.Println(sell.Person.Balance, sell.Amount)

		if sell.Person.Balance >= sell.Amount {
			sellLog = models.SellLog{
				User_id: sell.Person.ID,
				Amount:  sell.Amount,
				City:    sell.Address.City,
				State:   sell.Address.State,
				IsValid: true,
			}
		} else {
			sellLog = models.SellLog{
				User_id: sell.Person.ID,
				Amount:  sell.Amount,
				City:    sell.Address.City,
				State:   sell.Address.State,
				IsValid: false,
			}
		}

		insertLog(sellLog)
		updateBalance(sellLog.User_id, sell.Person.Balance-sell.Amount)
	} else {
		return
	}

}

func insertLog(s models.SellLog) {

	collection := models.Mongo.Database("public").Collection("transactions")

	document := bson.M{
		"user_id":  s.User_id,
		"amount":   s.Amount,
		"city":     s.City,
		"state":    s.State,
		"is_valid": s.IsValid,
	}

	_, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		log.Printf("Failed to insert new user: %s\n", err)
	}

}

func updateBalance(objID primitive.ObjectID, newBalance float64) {

	collection := models.Mongo.Database("public").Collection("users")

	update := bson.M{
		"$set": bson.M{
			"balance": newBalance,
		},
	}

	result, err := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		log.Fatalf("Error updating user balance: %v", err)
	}

	if result.MatchedCount == 0 {
		fmt.Println("No document matched the provided _id")
	} else {
		fmt.Printf("Matched %d document(s) and updated %d document(s)\n", result.MatchedCount, result.ModifiedCount)
	}

}
