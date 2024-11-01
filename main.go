package main

import (
	"fmt"

	"github.com/adamGrgic/riotgosdk/riotapis"
)

func main() {

	fmt.Println("Running test cycle for riot apis.")

	fmt.Println("Testing function dependency...")
	riotapis.TestFunction()

}
