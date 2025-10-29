import { useState, useMemo } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { createAppTheme } from './theme';
import Layout from './components/Layout';
import HomePage from './pages/HomePage';
import TopicPage from './pages/TopicPage';
import InvestmentsPage from './pages/InvestmentsPage';
import StockMarketsPage from './pages/StockMarketsPage';
import BondsPage from './pages/BondsPage';
import CryptocurrenciesPage from './pages/CryptocurrenciesPage';
import PersonalFinancePage from './pages/PersonalFinancePage';
import RealEstatePage from './pages/RealEstatePage';
import BankingPage from './pages/BankingPage';
import FintechPage from './pages/FintechPage';
import RiskManagementPage from './pages/RiskManagementPage';
import GlobalEconomyPage from './pages/GlobalEconomyPage';

function App() {
  const [mode, setMode] = useState<'light' | 'dark'>('light');

  const theme = useMemo(() => createAppTheme(mode), [mode]);

  const toggleColorMode = () => {
    setMode((prevMode) => (prevMode === 'light' ? 'dark' : 'light'));
  };

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Router>
        <Layout mode={mode} toggleColorMode={toggleColorMode}>
          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/topic/:id" element={<TopicPage />} />
            <Route path="/investments" element={<InvestmentsPage />} />
            <Route path="/stock-markets" element={<StockMarketsPage />} />
            <Route path="/bonds" element={<BondsPage />} />
            <Route path="/cryptocurrencies" element={<CryptocurrenciesPage />} />
            <Route path="/personal-finance" element={<PersonalFinancePage />} />
            <Route path="/real-estate" element={<RealEstatePage />} />
            <Route path="/banking" element={<BankingPage />} />
            <Route path="/fintech" element={<FintechPage />} />
            <Route path="/risk-management" element={<RiskManagementPage />} />
            <Route path="/global-economy" element={<GlobalEconomyPage />} />
          </Routes>
        </Layout>
      </Router>
    </ThemeProvider>
  );
}

export default App;
