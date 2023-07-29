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
			MarketCap:       100000000,
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
		{
			Name:            "Litecoin",
			Price:           100,
			Last24Variation: 0,
			MarketCap:       100000,
		},
		{
			Name:            "Monero",
			Price:           50,
			Last24Variation: 0,
			MarketCap:       10000,
		},
		{
			Name:            "Dogecoin",
			Price:           0.5,
			Last24Variation: 0,
			MarketCap:       1000,
		},
		{
			Name:            "Dash",
			Price:           100,
			Last24Variation: 0,
			MarketCap:       100,
		},
		{
			Name:            "Zcash",
			Price:           1000,
			Last24Variation: 0,
			MarketCap:       10,
		},
		{
			Name:            "Tether",
			Price:           1,
			Last24Variation: 0,
			MarketCap:       1,
		},
		{
			Name:            "Bitcoin Cash",
			Price:           100,
			Last24Variation: 0,
			MarketCap:       1,
		},
		{
			Name:            "Bitcoin SV",
			Price:           100,
			Last24Variation: 0,
			MarketCap:       1,
		},
		{
			Name:            "Binance Coin",
			Price:           100,
			Last24Variation: 0,
			MarketCap:       1,
		},
		{
			Name:            "EOS",
			Price:           100,
			Last24Variation: 0,
			MarketCap:       1,
		},
		{
			Name:            "Tezos",
			Price:           100,
			Last24Variation: 0,
			MarketCap:       1,
		},
		{
			Name:            "Cardano",
			Price:           100,
			Last24Variation: 0,
			MarketCap:       1,
		},
	}
}
