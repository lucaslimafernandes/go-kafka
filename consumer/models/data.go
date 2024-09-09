package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Sell struct {
	PersonId string
	Amount   float64
	Address  RealAddress
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
	User_id string
	Amount  float64
	City    string
	State   string
	IsValid bool
}

type PersonBson struct {
	ID primitive.ObjectID `bson:"_id,omitempty"` // ID do MongoDB
	// ID          string      `bson:"_id,omitempty"` // ID do MongoDB
	Name        string      `bson:"name"`        // Nome do usuário
	CreditCard  string      `bson:"credit_card"` // Número do cartão de crédito
	Balance     float64     `bson:"balance"`     // Saldo
	Address     string      `bson:"address"`     // Endereço
	City        string      `bson:"city"`        // Cidade
	State       string      `bson:"state"`       // Estado
	PostalCode  string      `bson:"postal_code"` // CEP
	Coordinates Coordinates `bson:"coordinates"` // Coordenadas (latitude e longitude)
}

type Coordinates struct {
	Lat  float64 `bson:"lat"`  // Latitude
	Long float64 `bson:"long"` // Longitude
}
