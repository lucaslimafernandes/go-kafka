package models

import "github.com/go-faker/faker/v4"

type Person struct {
	Id         int
	Name       string            `faker:"name"`
	CreditCard string            `faker:"cc_number"`
	Address    faker.RealAddress `faker:"real_address"`
}

type Sell struct {
	PersonId int
	Amount   float64
	Address  faker.RealAddress
}

type MinMaxId struct {
	Min int
	Max int
}
