package models

// FinanceTopic represents a financial education topic
type FinanceTopic struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Summary     string   `json:"summary"`
	Keywords    []string `json:"keywords"`
	Resources   []string `json:"resources"`
}

// StockQuote represents stock market data
type StockQuote struct {
	Symbol        string  `json:"symbol"`
	Price         float64 `json:"price"`
	Change        float64 `json:"change"`
	ChangePercent float64 `json:"changePercent"`
	Volume        int64   `json:"volume"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Open          float64 `json:"open"`
	PreviousClose float64 `json:"previousClose"`
	LastUpdated   string  `json:"lastUpdated"`
}

// CryptoPrice represents cryptocurrency data
type CryptoPrice struct {
	ID                    string  `json:"id"`
	Symbol                string  `json:"symbol"`
	Name                  string  `json:"name"`
	CurrentPrice          float64 `json:"currentPrice"`
	MarketCap             float64 `json:"marketCap"`
	MarketCapRank         int     `json:"marketCapRank"`
	PriceChange24h        float64 `json:"priceChange24h"`
	PriceChangePercent24h float64 `json:"priceChangePercent24h"`
	High24h               float64 `json:"high24h"`
	Low24h                float64 `json:"low24h"`
	CirculatingSupply     float64 `json:"circulatingSupply"`
	LastUpdated           string  `json:"lastUpdated"`
}

// CurrencyRate represents exchange rate data
type CurrencyRate struct {
	FromCurrency string  `json:"fromCurrency"`
	ToCurrency   string  `json:"toCurrency"`
	Rate         float64 `json:"rate"`
	Bid          float64 `json:"bid"`
	Ask          float64 `json:"ask"`
	LastUpdated  string  `json:"lastUpdated"`
}

// TimeSeriesData represents historical financial data
type TimeSeriesData struct {
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}

// MarketOverview represents general market statistics
type MarketOverview struct {
	TotalMarketCap float64      `json:"totalMarketCap"`
	Volume24h      float64      `json:"volume24h"`
	TopGainers     []StockQuote `json:"topGainers"`
	TopLosers      []StockQuote `json:"topLosers"`
	MostActive     []StockQuote `json:"mostActive"`
}

// APIResponse is a generic response wrapper
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}
