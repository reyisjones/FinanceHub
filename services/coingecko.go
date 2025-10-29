package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"financehub/models"
)

// CoinGeckoService handles CoinGecko API calls
type CoinGeckoService struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewCoinGeckoService creates a new CoinGecko service
func NewCoinGeckoService() *CoinGeckoService {
	return &CoinGeckoService{
		BaseURL: "https://api.coingecko.com/api/v3",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetCryptoPrice retrieves cryptocurrency price data
func (s *CoinGeckoService) GetCryptoPrice(coinID string) (*models.CryptoPrice, error) {
	url := fmt.Sprintf("%s/coins/markets?vs_currency=usd&ids=%s&order=market_cap_desc&per_page=1&page=1&sparkline=false&price_change_percentage=24h",
		s.BaseURL, coinID)

	resp, err := s.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch crypto price: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("coin not found")
	}

	coin := result[0]

	return &models.CryptoPrice{
		ID:                    fmt.Sprintf("%v", coin["id"]),
		Symbol:                fmt.Sprintf("%v", coin["symbol"]),
		Name:                  fmt.Sprintf("%v", coin["name"]),
		CurrentPrice:          getFloat(coin["current_price"]),
		MarketCap:             getFloat(coin["market_cap"]),
		MarketCapRank:         getInt(coin["market_cap_rank"]),
		PriceChange24h:        getFloat(coin["price_change_24h"]),
		PriceChangePercent24h: getFloat(coin["price_change_percentage_24h"]),
		High24h:               getFloat(coin["high_24h"]),
		Low24h:                getFloat(coin["low_24h"]),
		CirculatingSupply:     getFloat(coin["circulating_supply"]),
		LastUpdated:           time.Now().Format(time.RFC3339),
	}, nil
}

// GetTopCryptos retrieves top cryptocurrencies by market cap
func (s *CoinGeckoService) GetTopCryptos(limit int) ([]models.CryptoPrice, error) {
	url := fmt.Sprintf("%s/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=%d&page=1&sparkline=false&price_change_percentage=24h",
		s.BaseURL, limit)

	resp, err := s.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch top cryptos: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	var cryptos []models.CryptoPrice
	for _, coin := range result {
		cryptos = append(cryptos, models.CryptoPrice{
			ID:                    fmt.Sprintf("%v", coin["id"]),
			Symbol:                fmt.Sprintf("%v", coin["symbol"]),
			Name:                  fmt.Sprintf("%v", coin["name"]),
			CurrentPrice:          getFloat(coin["current_price"]),
			MarketCap:             getFloat(coin["market_cap"]),
			MarketCapRank:         getInt(coin["market_cap_rank"]),
			PriceChange24h:        getFloat(coin["price_change_24h"]),
			PriceChangePercent24h: getFloat(coin["price_change_percentage_24h"]),
			High24h:               getFloat(coin["high_24h"]),
			Low24h:                getFloat(coin["low_24h"]),
			CirculatingSupply:     getFloat(coin["circulating_supply"]),
			LastUpdated:           time.Now().Format(time.RFC3339),
		})
	}

	return cryptos, nil
}

// Helper functions to safely extract values
func getFloat(val interface{}) float64 {
	if val == nil {
		return 0
	}
	switch v := val.(type) {
	case float64:
		return v
	case int:
		return float64(v)
	default:
		return 0
	}
}

func getInt(val interface{}) int {
	if val == nil {
		return 0
	}
	switch v := val.(type) {
	case float64:
		return int(v)
	case int:
		return v
	default:
		return 0
	}
}
