# Description
Simple Go program to read a CSV file regarding financial transactions over a time period from the ING bank and insert into a database.

## Structure
Repo structure inspired from [here](https://medium.com/evendyne/getting-started-with-go-project-structure-ab8814ded9c3).

## Install
Run `go install`.

## Run
### DB 
To build:
`sudo docker build --progress=plain . -t gofinancedb`

To run:
`sudo docker run -p 5432:5432 -e POSTGRES_PASSWORD="atreides" --detach gofinancedb`

### Go code
Download ING csv from Mijn ING and put it in your `${HOME}` directory. Next do:

`go run . ~/ING_transactions_2024.csv` if you named your file `ING_transactions_2024.csv`.

Be careful with your personal data :). 

## Test
`go test gofinance/banktransaction`

## Steps
1. Label data 
