package banktransaction;

import (
	"log"
	"strconv"
	"strings"
	"fmt"
	"time"
	"gitlab.com/metakeule/fmtdate"
	"os"
)

type BankTransaction struct {
	date time.Time 
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
		} else {
			os.Exit(1)
		}
    } 

	return transactions
}

func fromCsvSingle(csvLine []string) *BankTransaction {
	creditType := transactionTypeFromString(csvLine[5]) // Credit or Debit

	if (creditType != UndefinedTransactionType) {
		
		transactionDate, err := parseDate(csvLine[0])
		if (err == nil) {
			return &BankTransaction{
				date: transactionDate.AddDate(0, 0, 1), // Random add a day 
				description: csvLine[1], 
				account: csvLine[2], 
				amount: parseAmount(csvLine[6], creditType), 
				remarks: csvLine[8],
			}
		}
	}

	return nil
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

func parseDate(transactionDate string) (time.Time, error) {
	return fmtdate.Parse("YYYYMMDD", transactionDate)
}