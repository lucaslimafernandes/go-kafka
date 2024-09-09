package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Sell struct {
	Person  PersonBson
	Amount  float64
	Address RealAddress
}

type RealAddress struct {
	Address     string
	City        string
	State       string
	PostalCode  string
	Coordinates Coordinates
}

type SellValidation struct {
	PersonId int
	Balance  float64
}

type SellLog struct {
	User_id primitive.ObjectID `bson:"_id,omitempty"`
	Amount  float64
	City    string
	State   string
	IsValid bool
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
