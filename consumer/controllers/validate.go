package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/lucaslimafernandes/go-kafka-cons/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Validate(msg string) {

	var sell models.Sell
	var sellLog models.SellLog
	json.Unmarshal([]byte(msg), &sell)

	userAcc := getUser(sell.PersonId)

	fmt.Println(sell)
	fmt.Println(userAcc)

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

}

func getUser(s string) models.PersonBson {

	collection := models.Mongo.Database("testdb").Collection("users")

	objectID, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		log.Printf("Invalid ID: %v", err)
	}

	var result models.PersonBson
	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No document found with the provided _id")
		} else {
			log.Fatal(err)
		}
	}

	return result

}

func insertLog(s models.SellLog, newBalance float64) {

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

	update := bson.M{
		"$set": bson.M{
			"balance": newBalance,
		},
	}

	objectID, _ := primitive.ObjectIDFromHex(s.User_id)
	result, err := collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, update)
	if err != nil {
		log.Fatalf("Error updating user balance: %v", err)
	}

	if result.MatchedCount == 0 {
		fmt.Println("No document matched the provided _id")
	} else {
		fmt.Printf("Matched %d document(s) and updated %d document(s)\n", result.MatchedCount, result.ModifiedCount)
	}

}
