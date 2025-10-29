import { Box, Container, Typography, Paper } from '@mui/material';

export default function InvestmentsPage() {
  return (
    <Container maxWidth="lg">
      <Box sx={{ my: 4 }}>
        <Typography variant="h3" component="h1" gutterBottom sx={{ fontWeight: 700 }}>
          Investments
        </Typography>
        <Paper sx={{ p: 3, mb: 4 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            About Investments
          </Typography>
          <Typography variant="body1" paragraph>
            Investment is the act of allocating resources, usually money, with the expectation of generating income or profit. 
            Understanding different investment types, risk levels, and strategies is crucial for building wealth over time. 
            Common investment vehicles include stocks, bonds, mutual funds, ETFs, and real estate.
          </Typography>
        </Paper>
        <Paper sx={{ p: 3 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            Learning Resources
          </Typography>
          <Typography component="div">
            <ul>
              <li><a href="https://www.investopedia.com/terms/i/investment.asp" target="_blank" rel="noopener noreferrer">Investopedia: Investment</a></li>
              <li><a href="https://www.sec.gov/investor/pubs/investorpubs.htm" target="_blank" rel="noopener noreferrer">SEC Investor Publications</a></li>
              <li><a href="https://www.investor.gov/introduction-investing" target="_blank" rel="noopener noreferrer">Investor.gov</a></li>
            </ul>
          </Typography>
        </Paper>
      </Box>
    </Container>
  );
}
