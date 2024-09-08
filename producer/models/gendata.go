package models

import (
	"context"
	"log"
	"math/rand"

	"github.com/go-faker/faker/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Selling() Sell {

	// mm := minMaxId()
	user := selectMongo()

	min := 1.0
	max := 1500.0
	sell := Sell{
		Person:  user,
		Amount:  min + rand.Float64()*(max-min),
		Address: faker.GetRealAddress(),
	}

	return sell

}

func selectMongo() PersonBson {

	collection := Mongo.Database("public").Collection("users")

	// pipeline := mongo.Pipeline{
	// 	{{"$sample", bson.D{{"size", 1}}}},
	// }

	pipeline := mongo.Pipeline{
		{
			{Key: "$sample", Value: bson.D{{Key: "size", Value: 1}}},
		},
	}

	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	var randomUser PersonBson
	if cursor.Next(context.Background()) {
		if err := cursor.Decode(&randomUser); err != nil {
			log.Fatal(err)
		}
	}

	return randomUser

}

func GeneratorData() (Person, error) {

	data := Person{}
	err := faker.FakeData(&data)
	if err != nil {
		return Person{}, err
	}

	return data, nil

}
