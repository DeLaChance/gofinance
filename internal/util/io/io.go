package io;

import (
    "encoding/csv"
	"os"
	"log"
)

func ReadCsv(csvFileName string) [][]string {

	file, err := os.Open(csvFileName)
	
	if err != nil { 
        log.Fatal("Error while reading the file", err) 
    }
	
	// Closes the file 
    defer file.Close() 

	reader := csv.NewReader(file) 
	records, err := reader.ReadAll() 

	if err != nil { 
        log.Println("Error reading records") 
    } 
      
	return records 
}