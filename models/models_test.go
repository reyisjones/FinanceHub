package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFinanceTopicStruct(t *testing.T) {
	topic := FinanceTopic{
		ID:          "1",
		Title:       "Test Topic",
		Description: "Test Description",
		Summary:     "Test Summary",
		Keywords:    []string{"Keyword1", "Keyword2"},
		Resources:   []string{"Resource1", "Resource2"},
	}

	assert.Equal(t, "1", topic.ID)
	assert.Equal(t, "Test Topic", topic.Title)
	assert.Equal(t, "Test Description", topic.Description)
	assert.Equal(t, "Test Summary", topic.Summary)
	assert.Len(t, topic.Keywords, 2)
	assert.Len(t, topic.Resources, 2)
}

func TestFinanceTopicValidation(t *testing.T) {
	tests := []struct {
		name    string
		topic   FinanceTopic
		isValid bool
	}{
		{
			name: "Valid topic",
			topic: FinanceTopic{
				ID:          "1",
				Title:       "Valid Topic",
				Description: "Valid Description",
				Summary:     "Valid Summary",
				Keywords:    []string{"Point"},
				Resources:   []string{"Resource"},
			},
			isValid: true,
		},
		{
			name: "Empty ID",
			topic: FinanceTopic{
				ID:          "",
				Title:       "Topic",
				Description: "Description",
				Summary:     "Summary",
				Keywords:    []string{"Point"},
				Resources:   []string{"Resource"},
			},
			isValid: false,
		},
		{
			name: "Empty title",
			topic: FinanceTopic{
				ID:          "1",
				Title:       "",
				Description: "Description",
				Summary:     "Summary",
				Keywords:    []string{"Point"},
				Resources:   []string{"Resource"},
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := tt.topic.ID != "" && tt.topic.Title != "" && tt.topic.Description != ""
			assert.Equal(t, tt.isValid, valid)
		})
	}
}

func TestStockQuoteStruct(t *testing.T) {
	quote := StockQuote{
		Symbol:        "AAPL",
		Price:         150.25,
		Change:        2.50,
		ChangePercent: 1.69,
		Volume:        1000000,
		High:          152.00,
		Low:           148.50,
		Open:          149.00,
	}

	assert.Equal(t, "AAPL", quote.Symbol)
	assert.Equal(t, 150.25, quote.Price)
	assert.Greater(t, quote.Volume, int64(0))
}

func TestCryptoPriceStruct(t *testing.T) {
	crypto := CryptoPrice{
		ID:                    "bitcoin",
		Symbol:                "BTC",
		Name:                  "Bitcoin",
		CurrentPrice:          45000.00,
		MarketCap:             850000000000,
		MarketCapRank:         1,
		PriceChange24h:        1200.00,
		PriceChangePercent24h: 2.74,
	}

	assert.Equal(t, "bitcoin", crypto.ID)
	assert.Equal(t, "BTC", crypto.Symbol)
	assert.Equal(t, 1, crypto.MarketCapRank)
	assert.Greater(t, crypto.MarketCap, float64(0))
}
