package models

import (
	"github.com/go-faker/faker/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `faker:"name"`
	CreditCard string             `faker:"cc_number"`
	Address    faker.RealAddress  `faker:"real_address"`
}

type Sell struct {
	Person  PersonBson
	Amount  float64
	Address faker.RealAddress
}

type MinMaxId struct {
	Min int
	Max int
}

type PersonBson struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	CreditCard  string             `bson:"credit_card"`
	Balance     float64            `bson:"balance"`
	Address     string             `bson:"address"`
	City        string             `bson:"city"`
	State       string             `bson:"state"`
	PostalCode  string             `bson:"postal_code"`
	Coordinates Coordinates        `bson:"coordinates"`
}

type Coordinates struct {
	Lat  float64 `bson:"lat"`
	Long float64 `bson:"long"`
}
