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
	ID          primitive.ObjectID `bson:"_id,omitempty"` // ID do MongoDB
	Name        string             `bson:"name"`          // Nome do usuário
	CreditCard  string             `bson:"credit_card"`   // Número do cartão de crédito
	Balance     float64            `bson:"balance"`       // Saldo
	Address     string             `bson:"address"`       // Endereço
	City        string             `bson:"city"`          // Cidade
	State       string             `bson:"state"`         // Estado
	PostalCode  string             `bson:"postal_code"`   // CEP
	Coordinates Coordinates        `bson:"coordinates"`   // Coordenadas (latitude e longitude)
}

type Coordinates struct {
	Lat  float64 `bson:"lat"`  // Latitude
	Long float64 `bson:"long"` // Longitude
}
