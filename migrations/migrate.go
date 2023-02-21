package main

import (
	"github.com/kkomissarov/beggar/db"
	"github.com/kkomissarov/beggar/initializers"
	"github.com/kkomissarov/beggar/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	db.ConnectToDatabase()
}

func main() {
	db.DB.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_types') THEN
				CREATE TYPE transaction_types AS ENUM (
					'income',
					'expense'
				);
			END IF;
		END$$;
	`)

	err := db.DB.AutoMigrate(
		&models.User{},
		&models.Currency{},
		&models.Account{},
		&models.TransactionCategory{},
		&models.Transaction{},
		&models.RevokedToken{},
	)
	if err != nil {
		log.Fatal("Unable to complete auto migrations")
	}
	log.Println("Migrations completed")
}
