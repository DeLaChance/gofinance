package banktransaction;

import (
    "testing"
	"time"
	"github.com/stretchr/testify/assert"
	"gitlab.com/metakeule/fmtdate"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestThatCsvLineCanBeMappedToTransaction(t *testing.T) {
	// Given
    csvLine := [9]string{
		"20240226",
		"BV",
		"NL58INGB********",
		"NL13ABNA0581483***",
		"IC",
		"Debit",
		"4,8",
		"SEPA direct debit",
		"Insurance bill",
	}
    
	// When
    actual := fromCsvSingle(csvLine[:])

	// Then
	expected := expectedTransaction()
	assert.Equal(t, actual, expected, "The two transactions should be the same.")
}

func expectedTransaction() *BankTransaction {
	return &BankTransaction{
		date: expectedDateTime(),
		description: "BV",
		account: "NL58INGB********",
		amount: -4.80,
		remarks: "Insurance bill",
	};
}

func expectedDateTime() time.Time {
	date, error := fmtdate.Parse("YYYYMMDD", "20240227")
	if (error != nil) {
		panic(error)
	}

	return date
}
