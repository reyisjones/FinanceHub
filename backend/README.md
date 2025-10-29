# FinanceHub Backend

Go backend server providing RESTful API endpoints for financial data.

## Quick Start

1. Install dependencies:
```bash
go mod download
```

2. Configure environment:
```bash
cp .env.example .env
# Edit .env and add your Alpha Vantage API key
```

3. Run the server:
```bash
go run main.go
```

Server will start on `http://localhost:8080`

## API Endpoints

- `GET /api/health` - Health check
- `GET /api/topics` - Get all finance topics
- `GET /api/topics/:id` - Get topic by ID
- `GET /api/stocks/:symbol` - Get stock quote
- `GET /api/stocks/:symbol/timeseries` - Get historical data
- `GET /api/crypto/top` - Get top cryptocurrencies
- `GET /api/crypto/:id` - Get crypto price
- `GET /api/currency/:from/:to` - Get exchange rate

## External APIs Used

- **Alpha Vantage**: Stock market data and currency rates
- **CoinGecko**: Cryptocurrency market data
