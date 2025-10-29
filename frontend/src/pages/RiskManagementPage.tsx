import { Box, Container, Typography, Paper } from '@mui/material';

export default function RiskManagementPage() {
  return (
    <Container maxWidth="lg">
      <Box sx={{ my: 4 }}>
        <Typography variant="h3" component="h1" gutterBottom sx={{ fontWeight: 700 }}>
          Risk Management
        </Typography>
        <Paper sx={{ p: 3, mb: 4 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            About Risk Management
          </Typography>
          <Typography variant="body1" paragraph>
            Risk management in finance involves identifying, analyzing, and mitigating potential losses in investments and financial decisions. 
            Strategies include diversification, hedging, insurance, and asset allocation. Understanding risk tolerance, risk-return tradeoff, and various types of risk 
            (market risk, credit risk, liquidity risk) is essential for protecting and growing wealth.
          </Typography>
        </Paper>
        <Paper sx={{ p: 3 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            Learning Resources
          </Typography>
          <Typography component="div">
            <ul>
              <li><a href="https://www.investopedia.com/terms/r/riskmanagement.asp" target="_blank" rel="noopener noreferrer">Investopedia: Risk Management</a></li>
              <li><a href="https://www.cfa.org/" target="_blank" rel="noopener noreferrer">CFA Institute</a></li>
              <li><a href="https://www.prmia.org/" target="_blank" rel="noopener noreferrer">PRMIA</a></li>
            </ul>
          </Typography>
        </Paper>
      </Box>
    </Container>
  );
}
