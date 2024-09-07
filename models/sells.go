package models

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/go-faker/faker/v4"
)

func Selling() (Sell, error) {

	mm := minMaxId()
	idUser := selectId(mm)

	sell := Sell{
		PersonId: idUser,
	}

	err := faker.FakeData(&sell)
	if err != nil {
		fmt.Println(err)
		return Sell{}, nil
	}

	return sell, nil

}

func GetUser(u int) SellValidation {

	var dataSell SellValidation

	query := `
		select
			u.id ,
			u.balance
		from users u
		where u.id = $1
		;
	`

	DB.QueryRow(context.Background(), query, u).Scan(
		&dataSell.PersonId,
		&dataSell.Balance,
	)

	return dataSell

}

func selectId(m MinMaxId) int {

	return rand.Intn(m.Max-m.Min) + m.Min

}

func minMaxId() MinMaxId {

	var mmData MinMaxId

	mmQuery := `
		select 
			min(u.id) ,
			max(u.id)
		from users u 
		;
	`
	DB.QueryRow(context.Background(), mmQuery).Scan(
		&mmData.Min,
		&mmData.Max,
	)

	return mmData

}
