import { Box, Container, Typography, Paper } from '@mui/material';

export default function BankingPage() {
  return (
    <Container maxWidth="lg">
      <Box sx={{ my: 4 }}>
        <Typography variant="h3" component="h1" gutterBottom sx={{ fontWeight: 700 }}>
          Banking
        </Typography>
        <Paper sx={{ p: 3, mb: 4 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            About Banking
          </Typography>
          <Typography variant="body1" paragraph>
            Banking institutions provide financial services including deposits, loans, credit, and payment processing. 
            Types of accounts include checking, savings, money market, and certificates of deposit (CDs). Banks are regulated to ensure stability and protect consumers. 
            Understanding banking products, fees, interest rates, and digital banking tools is important for managing finances effectively.
          </Typography>
        </Paper>
        <Paper sx={{ p: 3 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            Learning Resources
          </Typography>
          <Typography component="div">
            <ul>
              <li><a href="https://www.investopedia.com/terms/b/bank.asp" target="_blank" rel="noopener noreferrer">Investopedia: Banking</a></li>
              <li><a href="https://www.fdic.gov/" target="_blank" rel="noopener noreferrer">FDIC</a></li>
              <li><a href="https://www.bankrate.com/" target="_blank" rel="noopener noreferrer">Bankrate</a></li>
            </ul>
          </Typography>
        </Paper>
      </Box>
    </Container>
  );
}
