package models

import "fmt"

type Currency struct {
	FullName        string
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

func (c Currency) Area() float64 {
	return c.MarketCap
}

func (c Currency) DisplayName() string {
	return c.Name
}

func (c Currency) DisplayValue() string {
	return "$" + fmt.Sprintf("%.2f", c.Price)
}

func (c Currency) DisplaySubValue() string {
	return fmt.Sprintf("%.2f", c.Last24Variation) + "%"
}
