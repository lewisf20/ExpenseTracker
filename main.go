package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/lewisf20/ExpenseTracker/config"
	"github.com/lewisf20/ExpenseTracker/internal/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := config.GetDatabase()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", db)

	server.Start("8080")
}
