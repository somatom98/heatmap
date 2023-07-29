package models

type Currency struct {
	Name            string
	Price           float64
	Last24Variation float64
	MarketCap       float64
}

func (c Currency) Color() string {
	if c.Last24Variation < 0 {
		return "red"
	}
	return "green"
}
