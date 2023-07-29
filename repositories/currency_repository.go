package repositories

import (
	"github.com/somatom98/heatmap/models"
)

type CurrencyRepository struct {
}

func NewCurrencyRepository() *CurrencyRepository {
	return &CurrencyRepository{}
}

func (c *CurrencyRepository) GetAll() []models.Currency {
	return []models.Currency{
		{
			Name:            "Bitcoin",
			Price:           30000,
			Last24Variation: 12.5,
			MarketCap:       1000000000,
		},
		{
			Name:            "Ethereum",
			Price:           1500,
			Last24Variation: -33.3,
			MarketCap:       50000000,
		},
		{
			Name:            "Ripple",
			Price:           1.5,
			Last24Variation: 5,
			MarketCap:       1000000,
		},
	}
}
