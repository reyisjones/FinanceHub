import { useEffect, useState } from 'react';
import {
  Box,
  Container,
  Typography,
  Paper,
  Grid,
  Card,
  CardContent,
  CircularProgress,
  Alert,
  Chip,
  TextField,
  Button,
} from '@mui/material';
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer } from 'recharts';
import TrendingUpIcon from '@mui/icons-material/TrendingUp';
import TrendingDownIcon from '@mui/icons-material/TrendingDown';
import { financeAPI } from '../services/api';
import type { StockQuote, TimeSeriesData } from '../types';

export default function StockMarketsPage() {
  const [stockSymbol, setStockSymbol] = useState('AAPL');
  const [quote, setQuote] = useState<StockQuote | null>(null);
  const [timeSeries, setTimeSeries] = useState<TimeSeriesData[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchStockData = async (symbol: string) => {
    setLoading(true);
    setError(null);
    try {
      const [quoteData, timeSeriesData] = await Promise.all([
        financeAPI.getStockQuote(symbol),
        financeAPI.getStockTimeSeries(symbol),
      ]);
      setQuote(quoteData);
      setTimeSeries(timeSeriesData.reverse());
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch stock data');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchStockData(stockSymbol);
  }, []);

  const handleSearch = () => {
    if (stockSymbol.trim()) {
      fetchStockData(stockSymbol.toUpperCase());
    }
  };

  return (
    <Container maxWidth="lg">
      <Box sx={{ my: 4 }}>
        <Typography variant="h3" component="h1" gutterBottom sx={{ fontWeight: 700 }}>
          Stock Markets
        </Typography>

        <Paper sx={{ p: 3, mb: 4 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            About Stock Markets
          </Typography>
          <Typography variant="body1" paragraph>
            Stock markets are venues where buyers and sellers meet to exchange equity shares of public corporations.
            Major exchanges include NYSE, NASDAQ, and international markets. Stock prices fluctuate based on supply
            and demand, company performance, economic indicators, and investor sentiment.
          </Typography>
          <Typography variant="body1" paragraph>
            Understanding market mechanics, indices, and trading strategies is essential for investors. This page
            provides real-time stock quotes and historical price charts to help you track market movements.
          </Typography>
        </Paper>

        <Paper sx={{ p: 3, mb: 4 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            Real-Time Stock Data
          </Typography>

          <Box sx={{ display: 'flex', gap: 2, mb: 3 }}>
            <TextField
              label="Stock Symbol"
              variant="outlined"
              value={stockSymbol}
              onChange={(e) => setStockSymbol(e.target.value.toUpperCase())}
              onKeyPress={(e) => e.key === 'Enter' && handleSearch()}
              placeholder="e.g., AAPL, GOOGL, MSFT"
            />
            <Button variant="contained" onClick={handleSearch} disabled={loading}>
              Search
            </Button>
          </Box>

          {error && <Alert severity="error" sx={{ mb: 2 }}>{error}</Alert>}

          {loading ? (
            <Box sx={{ display: 'flex', justifyContent: 'center', py: 4 }}>
              <CircularProgress />
            </Box>
          ) : quote ? (
            <>
              <Grid container spacing={2} sx={{ mb: 4 }}>
                <Grid item xs={12} md={6}>
                  <Card>
                    <CardContent>
                      <Typography variant="h6" gutterBottom>
                        {quote.symbol}
                      </Typography>
                      <Typography variant="h3" sx={{ fontWeight: 700, mb: 1 }}>
                        ${quote.price.toFixed(2)}
                      </Typography>
                      <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                        {quote.change >= 0 ? (
                          <TrendingUpIcon color="success" />
                        ) : (
                          <TrendingDownIcon color="error" />
                        )}
                        <Chip
                          label={`${quote.change >= 0 ? '+' : ''}${quote.change.toFixed(2)} (${quote.changePercent.toFixed(2)}%)`}
                          color={quote.change >= 0 ? 'success' : 'error'}
                        />
                      </Box>
                    </CardContent>
                  </Card>
                </Grid>
                <Grid item xs={12} md={6}>
                  <Card>
                    <CardContent>
                      <Typography variant="h6" gutterBottom>
                        Market Details
                      </Typography>
                      <Box sx={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: 1 }}>
                        <Typography variant="body2" color="text.secondary">Open:</Typography>
                        <Typography variant="body2">${quote.open.toFixed(2)}</Typography>
                        <Typography variant="body2" color="text.secondary">High:</Typography>
                        <Typography variant="body2">${quote.high.toFixed(2)}</Typography>
                        <Typography variant="body2" color="text.secondary">Low:</Typography>
                        <Typography variant="body2">${quote.low.toFixed(2)}</Typography>
                        <Typography variant="body2" color="text.secondary">Volume:</Typography>
                        <Typography variant="body2">{quote.volume.toLocaleString()}</Typography>
                      </Box>
                    </CardContent>
                  </Card>
                </Grid>
              </Grid>

              {timeSeries.length > 0 && (
                <Box>
                  <Typography variant="h6" gutterBottom>
                    30-Day Price Chart
                  </Typography>
                  <ResponsiveContainer width="100%" height={400}>
                    <LineChart data={timeSeries}>
                      <CartesianGrid strokeDasharray="3 3" />
                      <XAxis dataKey="date" />
                      <YAxis domain={['auto', 'auto']} />
                      <Tooltip />
                      <Legend />
                      <Line type="monotone" dataKey="close" stroke="#1976d2" name="Close Price" />
                    </LineChart>
                  </ResponsiveContainer>
                </Box>
              )}
            </>
          ) : null}
        </Paper>

        <Paper sx={{ p: 3 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            Learning Resources
          </Typography>
          <Typography component="div">
            <ul>
              <li><a href="https://www.investopedia.com/terms/s/stockmarket.asp" target="_blank" rel="noopener noreferrer">Investopedia: Stock Market</a></li>
              <li><a href="https://www.sec.gov/fast-answers/answersmrktstruchtm.html" target="_blank" rel="noopener noreferrer">SEC: Market Structure</a></li>
              <li><a href="https://finance.yahoo.com/" target="_blank" rel="noopener noreferrer">Yahoo Finance</a></li>
            </ul>
          </Typography>
        </Paper>
      </Box>
    </Container>
  );
}
