package main

import (
	"os"
	"log"
	"gofinance/util/io"
	"gofinance/banktransaction"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Needs exactly 1 argument")
	}

	csvFileName := os.Args[1]
	
	log.Println("Reading from CSV file: ", csvFileName)
	csvLines := io.ReadCsv(csvFileName)

	transactions := banktransaction.FromCsv(csvLines)
	banktransaction.InsertIntoDatabase(transactions)	
}