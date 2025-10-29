# FinanceHub - Finance Learning & Insights Web Application

A comprehensive full-stack web application for financial education and real-time market insights. Built with **Go (Golang)** backend and **Material UI + Vite + TypeScript** frontend.

## 🌟 Features

- **📚 10 Key Finance Topics**: Comprehensive educational content covering:
  - Investments
  - Stock Markets
  - Bonds
  - Cryptocurrencies
  - Personal Finance
  - Real Estate
  - Banking
  - Fintech
  - Risk Management
  - Global Economy

![alt text](image.png)

- **📊 Real-Time Market Data**: 
  - Live stock quotes and historical charts
  - Top 10 cryptocurrencies by market cap
  - Currency exchange rates
  - Interactive data visualizations

- **🎨 Modern UI/UX**:
  - Material UI design system
  - Dark/Light mode toggle
  - Responsive design for all devices
  - Clean sidebar navigation

- **🚀 High Performance**:
  - Go backend for fast API responses
  - Vite for lightning-fast frontend development
  - TypeScript for type safety

## 📁 Project Structure

```
FinanceHub/
├── backend/                 # Go backend
│   ├── handlers/           # API route handlers
│   ├── models/             # Data models
│   ├── services/           # External API services
│   ├── main.go             # Application entry point
│   ├── go.mod              # Go module dependencies
│   └── .env.example        # Environment variables template
│
├── frontend/               # React + TypeScript frontend
│   ├── src/
│   │   ├── components/    # Reusable UI components
│   │   ├── pages/         # Page components
│   │   ├── services/      # API client
│   │   ├── types/         # TypeScript types
│   │   ├── App.tsx        # Main app component
│   │   ├── main.tsx       # Application entry
│   │   └── theme.ts       # Material UI theme
│   ├── package.json       # npm dependencies
│   ├── vite.config.ts     # Vite configuration
│   └── tsconfig.json      # TypeScript configuration
│
└── README.md              # This file
```

## 🛠️ Technology Stack

### Backend
- **Go 1.21+**: High-performance backend language
- **Gin**: Web framework for routing and middleware
- **CORS**: Cross-origin resource sharing
- **godotenv**: Environment variable management

### Frontend
- **React 18**: UI library
- **TypeScript**: Type-safe JavaScript
- **Material UI (MUI)**: Component library
- **Vite**: Build tool and dev server
- **React Router**: Client-side routing
- **Recharts**: Data visualization
- **Axios**: HTTP client

### External APIs
- **Alpha Vantage**: Stock market data and currency rates
- **CoinGecko**: Cryptocurrency prices and market data

## 🚀 Getting Started

### Prerequisites

- **Go 1.21 or higher**: [Download Go](https://golang.org/dl/)
- **Node.js 18+ and npm**: [Download Node.js](https://nodejs.org/)
- **Alpha Vantage API Key**: [Get Free API Key](https://www.alphavantage.co/support/#api-key)

### Backend Setup

1. **Navigate to backend directory**:
   ```bash
   cd backend
   ```

2. **Create environment file**:
   ```bash
   cp .env.example .env
   ```

3. **Configure API keys** in `.env`:
   ```env
   PORT=8080
   ALPHA_VANTAGE_API_KEY=your_actual_api_key_here
   ```

4. **Install Go dependencies**:
   ```bash
   go mod download
   ```

5. **Run the backend server**:
   ```bash
   go run main.go
   ```

   The API server will start on `http://localhost:8080`

### Frontend Setup

1. **Navigate to frontend directory**:
   ```bash
   cd frontend
   ```

2. **Install npm dependencies**:
   ```bash
   npm install
   ```

3. **Start the development server**:
   ```bash
   npm run dev
   ```

   The application will open at `http://localhost:5173`

## 🔑 API Key Configuration

### Alpha Vantage API Key

1. Visit [Alpha Vantage](https://www.alphavantage.co/support/#api-key)
2. Enter your email to receive a free API key
3. Copy the API key
4. Add it to `backend/.env` file:
   ```
   ALPHA_VANTAGE_API_KEY=YOUR_KEY_HERE
   ```

**Note**: Free tier provides 25 API requests per day. For production use, consider upgrading.

### CoinGecko API

CoinGecko's public API is used for cryptocurrency data and doesn't require an API key for basic usage.

## 📡 API Endpoints

### Topics
- `GET /api/topics` - Get all finance topics
- `GET /api/topics/:id` - Get specific topic by ID

### Stocks
- `GET /api/stocks/:symbol` - Get real-time stock quote
- `GET /api/stocks/:symbol/timeseries` - Get historical stock data (30 days)

### Cryptocurrencies
- `GET /api/crypto/top` - Get top 10 cryptocurrencies
- `GET /api/crypto/:id` - Get specific cryptocurrency price

### Currency Exchange
- `GET /api/currency/:from/:to` - Get exchange rate between currencies

### Health Check
- `GET /api/health` - API health status

## 🎨 Frontend Pages

- **/** - Homepage with topic overview
- **/investments** - Investments education page
- **/stock-markets** - Stock markets with live data and charts
- **/bonds** - Bonds education page
- **/cryptocurrencies** - Crypto prices and market data
- **/personal-finance** - Personal finance guidance
- **/real-estate** - Real estate investment information
- **/banking** - Banking systems overview
- **/fintech** - Financial technology insights
- **/risk-management** - Risk management strategies
- **/global-economy** - Global economy trends

## 🔧 Development

### Backend Development

```bash
cd backend
go run main.go
```

### Frontend Development

```bash
cd frontend
npm run dev
```

### Build for Production

**Backend**:
```bash
cd backend
go build -o financehub
./financehub
```

**Frontend**:
```bash
cd frontend
npm run build
npm run preview
```

## 🧪 Example Stock Symbols

Try these symbols in the Stock Markets page:
- **AAPL** - Apple Inc.
- **GOOGL** - Alphabet Inc.
- **MSFT** - Microsoft Corporation
- **AMZN** - Amazon.com Inc.
- **TSLA** - Tesla Inc.
- **NVDA** - NVIDIA Corporation

## 🪙 Example Cryptocurrency IDs

Common crypto IDs for the API:
- **bitcoin** - Bitcoin
- **ethereum** - Ethereum
- **binancecoin** - Binance Coin
- **cardano** - Cardano
- **solana** - Solana

## 🌍 Currency Codes

Common currency codes:
- **USD** - US Dollar
- **EUR** - Euro
- **GBP** - British Pound
- **JPY** - Japanese Yen
- **CAD** - Canadian Dollar

## 🎯 Features Implementation

### ✅ Educational Content
Each finance topic includes:
- Comprehensive summary
- Key concepts and keywords
- Curated learning resources
- External links to trusted sources

### ✅ Live Data Integration
- Real-time stock quotes from Alpha Vantage
- Historical price charts using Recharts
- Top cryptocurrencies from CoinGecko
- Currency exchange rates

### ✅ Responsive Design
- Mobile-first approach
- Adaptive layouts for all screen sizes
- Touch-friendly navigation
- Optimized for tablets and desktops

### ✅ Dark/Light Mode
- Toggle between themes
- Persistent user preference
- Material UI theming system

## 📝 Notes

- **API Rate Limits**: Be mindful of Alpha Vantage's rate limits (25 requests/day for free tier)
- **Error Handling**: The app includes error handling for failed API requests
- **Loading States**: Displays loading indicators while fetching data
- **Type Safety**: Full TypeScript coverage on frontend for better development experience

## 🚧 Future Enhancements

Potential features for future development:
- User authentication and personalized watchlists
- News feed integration
- Advanced charting with technical indicators
- Portfolio tracking
- Price alerts and notifications
- Mobile app version
- More data sources and APIs

## 📄 License

This project is open source and available for educational purposes.

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!

## 📧 Support

For questions or support, please open an issue in the repository.

---

**Built with ❤️ using Go and Material UI**
