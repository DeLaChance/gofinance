package main

import (
	"os"
	"log"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Needs exactly 1 argument")
	}

	csvFile := os.Args[1]
	log.Println("Reading from CSV file: ", csvFile)
}