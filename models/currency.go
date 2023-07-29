package models

type Currency struct {
	Name            string
	Price           float64
	Last24Variation float64
	MarketCap       float64
}
