package models

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {

	var err error

	var_db := "user=postgres password=password host=localhost port=5432 dbname=postgres"

	DB, err = pgxpool.New(context.Background(), var_db)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %s\n", err)
	}

}

func createTables() {

	sql_users := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			credit_card VARCHAR(20) NOT NULL,
			balance DECIMAL(10, 2) NOT NULL default 0,
			address VARCHAR(255),
			city VARCHAR(100),
			state VARCHAR(50),
			postal_code VARCHAR(20),
			coordinates_lat DECIMAL(9, 6),
			coordinates_long DECIMAL(9, 6)
		);
	`

	sql_logs := `
		CREATE TABLE IF NOT EXISTS logs (
			id SERIAL PRIMARY KEY,
			user_id INT ,
			amount DECIMAL(10, 2) NOT NULL,
			city VARCHAR(100),
			state VARCHAR(50),
			is_valid BOOLEAN NOT NULL,
			datetime TIMESTAMP DEFAULT now()
		);
	`

	_, err := DB.Exec(context.Background(), sql_users)
	if err != nil {
		log.Fatalf("Failed to create users table: %s\n", err)
	}

	_, err = DB.Exec(context.Background(), sql_logs)
	if err != nil {
		log.Fatalf("Failed to create logs table: %s\n", err)
	}
}

func Migrate() {

	createTables()

}
