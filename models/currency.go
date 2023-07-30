package models

import (
	"fmt"
	"math"
)

type Currency struct {
	FullName        string
	Name            string
	Price           float64
	Last24Variation float64
	MarketCap       float64
}

func (c Currency) Color() string {
	color := "green"
	if c.Last24Variation < 0 {
		color = "red"
	}

	if math.Abs(c.Last24Variation) < 10 {
		color = "dark" + color
	}

	return color
}

func (c Currency) Area() float64 {
	return c.MarketCap
}

func (c Currency) DisplayName() string {
	return c.Name
}

func (c Currency) DisplayValue() string {
	price := "$" + fmt.Sprintf("%.2f", c.Price)
	variation := fmt.Sprintf("%.2f", c.Last24Variation) + "%"
	if c.Last24Variation > 0 {
		variation = "+" + variation
	}
	return price + " (" + variation + ")"
}

func (c Currency) DisplaySubValue() string {
	return "$" + fmt.Sprintf("%.2f", c.MarketCap)
}
