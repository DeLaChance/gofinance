package banktransaction;

import (
	"log"
	"strconv"
	"strings"
	"fmt"
)

type BankTransaction struct {
	date string 
	description string
	account string 
	amount float64 // May be positive or negative
	remarks string
}

func (this BankTransaction) ToString() string {
	return fmt.Sprintf("description=%s, amount=%f, date=%s, remarks=%s, acount=%s", this.description, this.amount, this.date, this.remarks, this.account);
}

func FromCsv(csvLines [][]string) []BankTransaction {
	
	transactions := make([]BankTransaction, len(csvLines))

    for index, csvLine := range csvLines { 
		if (index == 0) {
			continue;
		}

        log.Printf("Processing line %d - %s", index, csvLine) 
		transaction := fromCsvSingle(csvLine)
		transactions = append(transactions, *transaction)
    } 

	return transactions
}

func fromCsvSingle(csvLine []string) *BankTransaction {
	return &BankTransaction{
		date: csvLine[0], // TODO: cast to date object 
		description: csvLine[1], 
		account: csvLine[2], 
		amount: parseAmount(csvLine[6]), 
		remarks: csvLine[8],
	}
}

func parseAmount(stringValue string) float64 {
	// TODO: cast to negative if deduction
	stringValue = strings.ReplaceAll(stringValue, ",", ".")

	floatValue, err := strconv.ParseFloat(stringValue, 2)
	if err != nil {
		log.Fatal("Invalid non-number value for: ", stringValue)
	}
	
	return floatValue
}