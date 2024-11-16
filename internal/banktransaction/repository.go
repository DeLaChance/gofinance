package banktransaction;

import (
	"database/sql"
	"fmt"
	"strings"
	"log"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "paul"
	password = "atreides"
	dbname   = "gofinance"
)

func InsertIntoDatabase(transactions []BankTransaction) error {
	db := connectDb()
	defer db.Close()

	sqlStatement := `	
		INSERT INTO gofinance.public.transactions ("date", description, account, amount, remarks) 
		VALUES %s;`

	valueStrings := make([]string, 0)
	valueArgs := []interface{}{}

	i := 0
	for _, t := range transactions {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", i*5+1, i*5+2, i*5+3, i*5+4, i*5+5))
 		
		valueArgs = append(valueArgs, t.date)
		valueArgs = append(valueArgs, t.description)
		valueArgs = append(valueArgs, t.account)
		valueArgs = append(valueArgs, t.amount)

		lengthOfRemarks := min(len(t.remarks), 128)
		valueArgs = append(valueArgs, t.remarks[:lengthOfRemarks])

		i++
	}

	sqlStatement = fmt.Sprintf(sqlStatement, strings.Join(valueStrings, ",\n"))
	
	tx, _ := db.Begin()
	_, execError := tx.Exec(sqlStatement, valueArgs...)
	if (execError != nil) {
		log.Fatal(execError)
		tx.Rollback()
		return execError
	}

	log.Printf("Inserted transactions successfully")
	
	return tx.Commit()
}

func connectDb() *sql.DB {
	
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    	"password=%s dbname=%s sslmode=disable",
    	host, port, user, password, dbname)

	log.Printf("Connecting to Postgres DB with details: %s", psqlInfo)		

	db, err := sql.Open("postgres", psqlInfo)
	if (err != nil) {
		 panic(err)
	}

	log.Printf("Connected to Postgres DB successfully...")		

	return db
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}