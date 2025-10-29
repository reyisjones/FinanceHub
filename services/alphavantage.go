package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"financehub/models"
)

// AlphaVantageService handles Alpha Vantage API calls
type AlphaVantageService struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

// NewAlphaVantageService creates a new Alpha Vantage service
func NewAlphaVantageService() *AlphaVantageService {
	return &AlphaVantageService{
		APIKey:  os.Getenv("ALPHA_VANTAGE_API_KEY"),
		BaseURL: "https://www.alphavantage.co/query",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetStockQuote retrieves real-time stock quote
func (s *AlphaVantageService) GetStockQuote(symbol string) (*models.StockQuote, error) {
	url := fmt.Sprintf("%s?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", s.BaseURL, symbol, s.APIKey)

	resp, err := s.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch stock quote: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	globalQuote, ok := result["Global Quote"].(map[string]interface{})
	if !ok || len(globalQuote) == 0 {
		return nil, fmt.Errorf("invalid response or symbol not found")
	}

	price, _ := strconv.ParseFloat(fmt.Sprintf("%v", globalQuote["05. price"]), 64)
	change, _ := strconv.ParseFloat(fmt.Sprintf("%v", globalQuote["09. change"]), 64)
	changePercent := fmt.Sprintf("%v", globalQuote["10. change percent"])
	changePercentVal, _ := strconv.ParseFloat(changePercent[:len(changePercent)-1], 64)
	volume, _ := strconv.ParseInt(fmt.Sprintf("%v", globalQuote["06. volume"]), 10, 64)
	high, _ := strconv.ParseFloat(fmt.Sprintf("%v", globalQuote["03. high"]), 64)
	low, _ := strconv.ParseFloat(fmt.Sprintf("%v", globalQuote["04. low"]), 64)
	open, _ := strconv.ParseFloat(fmt.Sprintf("%v", globalQuote["02. open"]), 64)
	prevClose, _ := strconv.ParseFloat(fmt.Sprintf("%v", globalQuote["08. previous close"]), 64)

	return &models.StockQuote{
		Symbol:        symbol,
		Price:         price,
		Change:        change,
		ChangePercent: changePercentVal,
		Volume:        volume,
		High:          high,
		Low:           low,
		Open:          open,
		PreviousClose: prevClose,
		LastUpdated:   time.Now().Format(time.RFC3339),
	}, nil
}

// GetTimeSeriesDaily retrieves daily time series data
func (s *AlphaVantageService) GetTimeSeriesDaily(symbol string, limit int) ([]models.TimeSeriesData, error) {
	url := fmt.Sprintf("%s?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s", s.BaseURL, symbol, s.APIKey)

	resp, err := s.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch time series: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	timeSeries, ok := result["Time Series (Daily)"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid time series data")
	}

	var data []models.TimeSeriesData
	count := 0
	for date, values := range timeSeries {
		if count >= limit {
			break
		}

		valueMap := values.(map[string]interface{})
		open, _ := strconv.ParseFloat(fmt.Sprintf("%v", valueMap["1. open"]), 64)
		high, _ := strconv.ParseFloat(fmt.Sprintf("%v", valueMap["2. high"]), 64)
		low, _ := strconv.ParseFloat(fmt.Sprintf("%v", valueMap["3. low"]), 64)
		close, _ := strconv.ParseFloat(fmt.Sprintf("%v", valueMap["4. close"]), 64)
		volume, _ := strconv.ParseInt(fmt.Sprintf("%v", valueMap["5. volume"]), 10, 64)

		data = append(data, models.TimeSeriesData{
			Date:   date,
			Open:   open,
			High:   high,
			Low:    low,
			Close:  close,
			Volume: volume,
		})
		count++
	}

	return data, nil
}

// GetCurrencyExchangeRate retrieves currency exchange rate
func (s *AlphaVantageService) GetCurrencyExchangeRate(fromCurrency, toCurrency string) (*models.CurrencyRate, error) {
	url := fmt.Sprintf("%s?function=CURRENCY_EXCHANGE_RATE&from_currency=%s&to_currency=%s&apikey=%s",
		s.BaseURL, fromCurrency, toCurrency, s.APIKey)

	resp, err := s.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch exchange rate: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	exchangeData, ok := result["Realtime Currency Exchange Rate"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid exchange rate data")
	}

	rate, _ := strconv.ParseFloat(fmt.Sprintf("%v", exchangeData["5. Exchange Rate"]), 64)
	bid, _ := strconv.ParseFloat(fmt.Sprintf("%v", exchangeData["8. Bid Price"]), 64)
	ask, _ := strconv.ParseFloat(fmt.Sprintf("%v", exchangeData["9. Ask Price"]), 64)

	return &models.CurrencyRate{
		FromCurrency: fromCurrency,
		ToCurrency:   toCurrency,
		Rate:         rate,
		Bid:          bid,
		Ask:          ask,
		LastUpdated:  time.Now().Format(time.RFC3339),
	}, nil
}
