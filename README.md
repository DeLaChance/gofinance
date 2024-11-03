# Description
Simple Go program to read a CSV file regarding financial transactions over a time period from the ING bank and insert into a database.

## Install
Run `go install`.

## Run
Download ING csv from Mijn ING and put it in your `${HOME}` directory. Next do:

`go run . ~/ING_transactions_2024.csv` if you named your file `ING_transactions_2024.csv`.

## Steps
1. Read a file at a CLI parameter
2. Write a class for Transaction & print file.
3. Connect to DB.
4. Insert into DB.