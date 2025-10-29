# FinanceHub - Setup Guide

## 🎯 Quick Setup (Step-by-Step)

### Step 1: Get API Key

1. Visit https://www.alphavantage.co/support/#api-key
2. Enter your email address
3. You'll receive an API key instantly (free tier: 25 requests/day)

### Step 2: Backend Setup

```bash
# Navigate to backend directory
cd c:\Dev\FinanceHub\backend

# Copy environment template
copy .env.example .env

# Edit .env file and replace with your actual API key
# ALPHA_VANTAGE_API_KEY=your_actual_key_here

# Install Go dependencies
go mod download

# Start the backend server
go run main.go
```

The backend will start at `http://localhost:8080`

### Step 3: Frontend Setup

Open a **new terminal window**:

```bash
# Navigate to frontend directory
cd c:\Dev\FinanceHub\frontend

# Install dependencies
npm install

# Start the development server
npm run dev
```

The frontend will open at `http://localhost:5173`

### Step 4: Access the Application

Open your browser and go to `http://localhost:5173`

## ✅ Verification Checklist

- [ ] Backend running on port 8080
- [ ] Frontend running on port 5173
- [ ] Can access homepage
- [ ] Can navigate to Stock Markets page
- [ ] Can search for stock symbols (try AAPL)
- [ ] Can see top cryptocurrencies
- [ ] Dark/Light mode toggle works

## 🐛 Troubleshooting

### Backend Issues

**Error: "Missing API key"**
- Solution: Make sure you've created `.env` file and added your Alpha Vantage API key

**Error: "Port 8080 already in use"**
- Solution: Change PORT in `.env` file to another port (e.g., 8081)
- Update frontend API_BASE_URL accordingly

**Error: "Cannot download Go modules"**
- Solution: Run `go mod tidy` then `go mod download`

### Frontend Issues

**Error: "Dependencies not found"**
- Solution: Delete `node_modules` folder and `package-lock.json`, then run `npm install` again

**Error: "Port 5173 already in use"**
- Solution: The dev server will automatically try another port (5174, etc.)

**Error: "Cannot connect to backend"**
- Solution: Make sure backend is running on port 8080
- Check CORS settings in backend `main.go`

### API Issues

**Error: "Invalid API key" or "API request failed"**
- Solution: Verify your Alpha Vantage API key is correct
- Check if you've exceeded the free tier limit (25 requests/day)

**Error: "Symbol not found"**
- Solution: Use valid stock symbols (e.g., AAPL, GOOGL, MSFT)

## 📊 Testing the Application

### Test Stock Market Features

1. Go to "Stock Markets" page
2. Search for: `AAPL`
3. You should see:
   - Current stock price
   - Price change
   - Market details (Open, High, Low, Volume)
   - 30-day price chart

### Test Cryptocurrency Features

1. Go to "Cryptocurrencies" page
2. You should see:
   - Top 10 cryptocurrencies by market cap
   - Current prices
   - 24h price changes
   - Market cap values

### Test Navigation

1. Click through all 10 finance topics in the sidebar
2. Each page should load with educational content
3. Dark/Light mode toggle should work on all pages

## 🚀 Development Tips

### Hot Reload

Both frontend and backend support hot reload:
- **Frontend**: Vite auto-reloads on file changes
- **Backend**: Use `air` for Go hot reload:
  ```bash
  go install github.com/cosmtrek/air@latest
  air
  ```

### Code Organization

- **Backend handlers**: Add new API routes in `backend/handlers/`
- **Frontend pages**: Add new pages in `frontend/src/pages/`
- **UI components**: Create reusable components in `frontend/src/components/`
- **API calls**: Add new API functions in `frontend/src/services/api.ts`

## 📁 Key Files Reference

### Backend
- `main.go` - Server entry point and routes
- `handlers/handlers.go` - API request handlers
- `services/*.go` - External API integrations
- `models/models.go` - Data structures
- `.env` - Environment variables (API keys)

### Frontend
- `src/main.tsx` - Application entry point
- `src/App.tsx` - Main app with routing
- `src/theme.ts` - Material UI theme configuration
- `src/services/api.ts` - Backend API client
- `src/types/index.ts` - TypeScript type definitions

## 🎨 Customization

### Change Theme Colors

Edit `frontend/src/theme.ts` to customize colors:

```typescript
primary: {
  main: '#1976d2',  // Change to your preferred color
}
```

### Add New Finance Topics

1. Add topic data in `backend/services/topics.go`
2. Create new page component in `frontend/src/pages/`
3. Add route in `frontend/src/App.tsx`
4. Add menu item in `frontend/src/components/Layout.tsx`

### Add New API Endpoints

1. Create handler function in `backend/handlers/handlers.go`
2. Add route in `backend/main.go`
3. Add API call in `frontend/src/services/api.ts`
4. Use in page components

## 📚 Learning Resources

- **Go**: https://go.dev/tour/
- **React**: https://react.dev/
- **Material UI**: https://mui.com/
- **TypeScript**: https://www.typescriptlang.org/docs/
- **Vite**: https://vitejs.dev/

## 🎉 You're All Set!

Your full-stack Finance Learning and Insights application is now running. Explore the features, check out the real-time data, and customize it to your needs!

For questions or issues, refer to the main README.md file.
