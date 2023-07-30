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
			FullName:        "Bitcoin",
			Name:            "BTC",
			Price:           30000,
			Last24Variation: 12.5,
			MarketCap:       1000000,
		},
		{
			FullName:        "Ethereum",
			Name:            "ETH",
			Price:           1500,
			Last24Variation: -33.3,
			MarketCap:       500000,
		},
		{
			FullName:        "Ripple",
			Name:            "XRP",
			Price:           1.5,
			Last24Variation: 5,
			MarketCap:       500000,
		},
		{
			FullName:        "Litecoin",
			Name:            "LTC",
			Price:           100,
			Last24Variation: -1,
			MarketCap:       200000,
		},
		{
			FullName:        "Monero",
			Name:            "XMR",
			Price:           50,
			Last24Variation: 11.4,
			MarketCap:       100000,
		},
		{
			FullName:        "Dogecoin",
			Name:            "DOGE",
			Price:           0.5,
			Last24Variation: 0,
			MarketCap:       50000,
		},
		{
			FullName:        "Dash",
			Name:            "DASH",
			Price:           100,
			Last24Variation: 5,
			MarketCap:       100,
		},
		{
			FullName:        "Zcash",
			Name:            "ZEC",
			Price:           1000,
			Last24Variation: -2.4,
			MarketCap:       10,
		},
		{
			FullName:        "Tether",
			Name:            "USDT",
			Price:           1,
			Last24Variation: 0,
			MarketCap:       1,
		},
		{
			FullName:        "Bitcoin Cash",
			Name:            "BCH",
			Price:           100,
			Last24Variation: 0,
			MarketCap:       1,
		},
		{
			FullName:        "Bitcoin SV",
			Name:            "BSV",
			Price:           100,
			Last24Variation: 0,
			MarketCap:       1,
		},
		{
			FullName:        "Binance Coin",
			Name:            "BNB",
			Price:           100,
			Last24Variation: 0,
			MarketCap:       1,
		},
		{
			FullName:        "EOS",
			Name:            "EOS",
			Price:           100,
			Last24Variation: 0,
			MarketCap:       1,
		},
		{
			FullName:        "Tezos",
			Name:            "XTZ",
			Price:           100,
			Last24Variation: 0,
			MarketCap:       1,
		},
		{
			FullName:        "Cardano",
			Name:            "ADA",
			Price:           100,
			Last24Variation: 0,
			MarketCap:       1,
		},
	}
}
