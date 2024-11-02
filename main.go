package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adamGrgic/riotgosdk/riotapis"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Running test cycle for riot apis.")

	fmt.Println("Testing function dependency...")
	riotapis.TestFunction()

	fmt.Println("Testing API Connection...")

	apiKey := os.Getenv("API_KEY")
	riotapis.GetLeagueEntries(apiKey)

}
