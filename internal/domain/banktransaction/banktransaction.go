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

type BankTransactionType string 

const (
	UndefinedTransactionType BankTransactionType = ""
	Credit BankTransactionType = "Credit"
	Debit BankTransactionType = "Debit"
)

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
		if (transaction != nil) {
			transactions = append(transactions, *transaction)
		}
    } 

	return transactions
}

func fromCsvSingle(csvLine []string) *BankTransaction {
	creditType := transactionTypeFromString(csvLine[5]) // Credit or Debit
	if (creditType != UndefinedTransactionType) {
		return &BankTransaction{
			date: csvLine[0], // TODO: cast to date object 
			description: csvLine[1], 
			account: csvLine[2], 
			amount: parseAmount(csvLine[6], creditType), 
			remarks: csvLine[8],
		}
	} else {
		return nil
	}
}

func parseAmount(amountStringValue string, creditType BankTransactionType) float64 {
	amountStringValue = strings.ReplaceAll(amountStringValue, ",", ".")

	floatValue, err := strconv.ParseFloat(amountStringValue, 2)
	if err != nil {
		log.Fatal("Invalid non-number value for: ", amountStringValue)
	}

	if (creditType == Debit) {
		floatValue = -1.0 * floatValue;
	}
	
	return floatValue
}

func transactionTypeFromString(transactionType string) BankTransactionType {
	if (transactionType == "Credit") {
		return Credit;
	} else if (transactionType == "Debit") {
		return Debit;
	} else {
		return UndefinedTransactionType;
	}
}