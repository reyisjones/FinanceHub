# FinanceHub - Quick Reference

## 🚀 Quick Start Commands

### First Time Setup
```bash
# Windows
setup.bat

# Manual Setup
cd backend && copy .env.example .env
# Edit .env and add your Alpha Vantage API key
cd backend && go mod download
cd frontend && npm install
```

### Start Development
```bash
# Windows (starts both servers)
start-dev.bat

# Manual Start - Terminal 1 (Backend)
cd backend
go run main.go

# Manual Start - Terminal 2 (Frontend)
cd frontend
npm run dev
```

## 🌐 URLs

| Service | URL | Description |
|---------|-----|-------------|
| Frontend | http://localhost:5173 | Web application |
| Backend | http://localhost:8080 | API server |
| API Health | http://localhost:8080/api/health | Health check |
| API Topics | http://localhost:8080/api/topics | Finance topics |

## 📡 API Quick Reference

### Stocks
```bash
# Get stock quote
GET http://localhost:8080/api/stocks/AAPL

# Get historical data
GET http://localhost:8080/api/stocks/AAPL/timeseries
```

### Crypto
```bash
# Top 10 cryptocurrencies
GET http://localhost:8080/api/crypto/top

# Specific crypto
GET http://localhost:8080/api/crypto/bitcoin
```

### Currency
```bash
# Exchange rate
GET http://localhost:8080/api/currency/USD/EUR
```

## 🎯 Test Data

### Stock Symbols
- AAPL - Apple
- GOOGL - Google
- MSFT - Microsoft
- AMZN - Amazon
- TSLA - Tesla

### Crypto IDs
- bitcoin
- ethereum
- binancecoin
- cardano
- solana

### Currency Codes
- USD, EUR, GBP, JPY, CAD, AUD, CHF, CNY

## 🔑 Environment Variables

### Backend (.env)
```env
PORT=8080
ALPHA_VANTAGE_API_KEY=your_key_here
```

### Frontend (api.ts)
```typescript
const API_BASE_URL = 'http://localhost:8080/api';
```

## 🛠️ Common Commands

### Backend
```bash
# Install dependencies
go mod download

# Run server
go run main.go

# Build binary
go build -o financehub

# Run binary
./financehub
```

### Frontend
```bash
# Install dependencies
npm install

# Start dev server
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview

# Run linter
npm run lint
```

## 🐛 Quick Troubleshooting

| Issue | Solution |
|-------|----------|
| Backend won't start | Check if .env exists and has valid API key |
| Port 8080 in use | Change PORT in .env |
| Frontend errors | Delete node_modules, run npm install |
| API rate limit | Wait (Alpha Vantage: 25/day free) |
| Stock not found | Use valid symbols (AAPL, GOOGL, etc.) |
| CORS errors | Check backend CORS config in main.go |

## 📁 Key Files

| File | Purpose |
|------|---------|
| backend/main.go | Server entry point |
| backend/.env | API keys and config |
| frontend/src/App.tsx | Main app and routes |
| frontend/src/services/api.ts | Backend API calls |
| frontend/src/theme.ts | UI theme config |

## 🎨 Navigation Structure

```
Home
├── Investments
├── Stock Markets (with live data)
├── Bonds
├── Cryptocurrencies (with live data)
├── Personal Finance
├── Real Estate
├── Banking
├── Fintech
├── Risk Management
└── Global Economy
```

## 💡 Pro Tips

1. **API Limits**: Free Alpha Vantage = 25 requests/day
2. **Dark Mode**: Toggle icon in top-right corner
3. **Mobile**: Hamburger menu appears on small screens
4. **Charts**: Hover over chart for details
5. **Search**: Try different stock symbols
6. **Reload**: Refresh page to fetch latest data

## 📞 Getting Help

- Main docs: `README.md`
- Setup guide: `SETUP.md`
- Project summary: `PROJECT_SUMMARY.md`
- Backend docs: `backend/README.md`
- Frontend docs: `frontend/README.md`

## ✅ Verification Checklist

- [ ] Go and Node.js installed
- [ ] Alpha Vantage API key obtained
- [ ] Backend .env configured
- [ ] Backend running (port 8080)
- [ ] Frontend running (port 5173)
- [ ] Can access homepage
- [ ] Can navigate all topics
- [ ] Stock search works
- [ ] Crypto list displays
- [ ] Dark mode toggles

---
**Version**: 1.0 | **Last Updated**: 2025
