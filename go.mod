module main

require gofinance/util/io v1.0.0

require gofinance/banktransaction v1.0.0

replace gofinance/util/io => ./internal/util/io

replace gofinance/banktransaction => ./internal/banktransaction

require gitlab.com/metakeule/fmtdate v1.2.2

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

go 1.18
