package main

import (
	"github.com/kkomissarov/beggar/db"
	"github.com/kkomissarov/beggar/initializers"
	"github.com/kkomissarov/beggar/router"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	db.ConnectToDatabase()
}

func main() {
	r := router.Router()
	if err := r.Run(":3000"); err != nil {
		log.Fatal("Unable to run server")
	}
}
