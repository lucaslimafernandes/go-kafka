package models

import "github.com/go-faker/faker/v4"

type Sell struct {
	PersonId int
	Amount   float64
	Address  faker.RealAddress
}

type MinMaxId struct {
	Min int
	Max int
}
