package models

type Stocks struct {
	ID     int
	Name   string
	Symbol string
	Price  []OHLC
}
