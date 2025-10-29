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
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from '@mui/material';
import TrendingUpIcon from '@mui/icons-material/TrendingUp';
import TrendingDownIcon from '@mui/icons-material/TrendingDown';
import { financeAPI } from '../services/api';
import type { CryptoPrice } from '../types';

export default function CryptocurrenciesPage() {
  const [topCryptos, setTopCryptos] = useState<CryptoPrice[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await financeAPI.getTopCryptos();
        setTopCryptos(data);
      } catch (err) {
        setError(err instanceof Error ? err.message : 'Failed to fetch crypto data');
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  return (
    <Container maxWidth="lg">
      <Box sx={{ my: 4 }}>
        <Typography variant="h3" component="h1" gutterBottom sx={{ fontWeight: 700 }}>
          Cryptocurrencies
        </Typography>

        <Paper sx={{ p: 3, mb: 4 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            About Cryptocurrencies
          </Typography>
          <Typography variant="body1" paragraph>
            Cryptocurrencies are digital or virtual currencies that use cryptography for security and operate on 
            decentralized networks based on blockchain technology. Bitcoin, Ethereum, and thousands of altcoins exist.
          </Typography>
          <Typography variant="body1" paragraph>
            Crypto offers potential for high returns but comes with significant volatility and risk. Key concepts include 
            blockchain, mining, wallets, decentralized finance (DeFi), and smart contracts.
          </Typography>
        </Paper>

        <Paper sx={{ p: 3, mb: 4 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            Top Cryptocurrencies by Market Cap
          </Typography>

          {error && <Alert severity="error" sx={{ mb: 2 }}>{error}</Alert>}

          {loading ? (
            <Box sx={{ display: 'flex', justifyContent: 'center', py: 4 }}>
              <CircularProgress />
            </Box>
          ) : (
            <>
              <Grid container spacing={2} sx={{ mb: 4 }}>
                {topCryptos.slice(0, 3).map((crypto) => (
                  <Grid item xs={12} md={4} key={crypto.id}>
                    <Card>
                      <CardContent>
                        <Typography variant="h6" gutterBottom>
                          {crypto.name} ({crypto.symbol.toUpperCase()})
                        </Typography>
                        <Typography variant="h4" sx={{ fontWeight: 700, mb: 1 }}>
                          ${crypto.currentPrice.toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 2 })}
                        </Typography>
                        <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                          {crypto.priceChangePercent24h >= 0 ? (
                            <TrendingUpIcon color="success" />
                          ) : (
                            <TrendingDownIcon color="error" />
                          )}
                          <Chip
                            label={`${crypto.priceChangePercent24h >= 0 ? '+' : ''}${crypto.priceChangePercent24h.toFixed(2)}%`}
                            color={crypto.priceChangePercent24h >= 0 ? 'success' : 'error'}
                            size="small"
                          />
                        </Box>
                        <Typography variant="body2" color="text.secondary" sx={{ mt: 1 }}>
                          Market Cap: ${(crypto.marketCap / 1e9).toFixed(2)}B
                        </Typography>
                      </CardContent>
                    </Card>
                  </Grid>
                ))}
              </Grid>

              <TableContainer>
                <Table>
                  <TableHead>
                    <TableRow>
                      <TableCell>#</TableCell>
                      <TableCell>Name</TableCell>
                      <TableCell align="right">Price</TableCell>
                      <TableCell align="right">24h Change</TableCell>
                      <TableCell align="right">Market Cap</TableCell>
                      <TableCell align="right">24h Volume</TableCell>
                    </TableRow>
                  </TableHead>
                  <TableBody>
                    {topCryptos.map((crypto) => (
                      <TableRow key={crypto.id}>
                        <TableCell>{crypto.marketCapRank}</TableCell>
                        <TableCell>
                          <Box>
                            <Typography variant="body1" fontWeight={600}>
                              {crypto.name}
                            </Typography>
                            <Typography variant="body2" color="text.secondary">
                              {crypto.symbol.toUpperCase()}
                            </Typography>
                          </Box>
                        </TableCell>
                        <TableCell align="right">
                          ${crypto.currentPrice.toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 2 })}
                        </TableCell>
                        <TableCell align="right">
                          <Chip
                            label={`${crypto.priceChangePercent24h >= 0 ? '+' : ''}${crypto.priceChangePercent24h.toFixed(2)}%`}
                            color={crypto.priceChangePercent24h >= 0 ? 'success' : 'error'}
                            size="small"
                          />
                        </TableCell>
                        <TableCell align="right">
                          ${(crypto.marketCap / 1e9).toFixed(2)}B
                        </TableCell>
                        <TableCell align="right">
                          ${(crypto.marketCap / 1e6).toFixed(0)}M
                        </TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </TableContainer>
            </>
          )}
        </Paper>

        <Paper sx={{ p: 3 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            Learning Resources
          </Typography>
          <Typography component="div">
            <ul>
              <li><a href="https://www.investopedia.com/terms/c/cryptocurrency.asp" target="_blank" rel="noopener noreferrer">Investopedia: Cryptocurrency</a></li>
              <li><a href="https://www.coingecko.com/" target="_blank" rel="noopener noreferrer">CoinGecko</a></li>
              <li><a href="https://ethereum.org/en/what-is-ethereum/" target="_blank" rel="noopener noreferrer">Ethereum Foundation</a></li>
            </ul>
          </Typography>
        </Paper>
      </Box>
    </Container>
  );
}
