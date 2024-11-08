module main

require gofinance/util/io v1.0.0
require gofinance/domain/banktransaction v1.0.0

replace gofinance/util/io => ./internal/util/io
replace gofinance/domain/banktransaction => ./internal/domain/banktransaction

require (
	"gitlab.com/metakeule/fmtdate" v1.2.2
)

go 1.18
