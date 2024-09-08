package models

import (
	"context"
	"math/rand"

	"github.com/go-faker/faker/v4"
)

func Selling() (Sell, error) {

	mm := minMaxId()
	idUser := selectId(mm)

	min := 1.0
	max := 1500.0
	sell := Sell{
		PersonId: idUser,
		Amount:   min + rand.Float64()*(max-min),
		Address:  faker.GetRealAddress(),
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
	// log.Print("\n\n\n", dataSell, "\n\n\n")

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
