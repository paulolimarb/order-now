package models

type Order struct {
}

type product struct {
	Name           string
	Specifications string
	PriceInCents   int64
}

type buyer struct {
	Name           string
	DocumentNumber string
	Address        string
}
