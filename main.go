package main

import (
	"os"
	"log"
	"internal/io"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Needs exactly 1 argument")
	}

	csvFileName := os.Args[1]
	log.Println("Reading from CSV file: ", csvFileName)
	io.ReadCsv(csvFileName)
}