package models

import "github.com/go-faker/faker/v4"

type Person struct {
	Id         int
	Name       string            `faker:"name"`
	CreditCard string            `faker:"cc_number"`
	Address    faker.RealAddress `faker:"real_address"`
}

type MinMaxId struct {
	Min int
	Max int
}

type Sell struct {
	PersonId int
	Amount   float64
	Address  faker.RealAddress
}

type SellValidation struct {
	PersonId int
	Balance  float64
}

type SellLog struct {
	User_id int
	Amount  float64
	City    string
	State   string
	IsValid bool
}

type ResponseSell struct {
	PersonId int
	Amount   float64
	Address  faker.RealAddress
	IsValid  bool
}

func GeneratorData() (Person, error) {

	data := Person{}
	err := faker.FakeData(&data)
	if err != nil {
		return Person{}, err
	}

	return data, nil

}
