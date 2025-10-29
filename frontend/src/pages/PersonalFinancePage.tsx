import { Box, Container, Typography, Paper } from '@mui/material';

export default function PersonalFinancePage() {
  return (
    <Container maxWidth="lg">
      <Box sx={{ my: 4 }}>
        <Typography variant="h3" component="h1" gutterBottom sx={{ fontWeight: 700 }}>
          Personal Finance
        </Typography>
        <Paper sx={{ p: 3, mb: 4 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            About Personal Finance
          </Typography>
          <Typography variant="body1" paragraph>
            Personal finance encompasses budgeting, saving, investing, insurance, mortgages, retirement planning, and tax planning. 
            Key principles include living below your means, building an emergency fund, eliminating high-interest debt, and saving for retirement. 
            Financial literacy helps individuals make informed decisions about spending, saving, and investing for long-term financial security.
          </Typography>
        </Paper>
        <Paper sx={{ p: 3 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            Learning Resources
          </Typography>
          <Typography component="div">
            <ul>
              <li><a href="https://www.investopedia.com/personal-finance-4427760" target="_blank" rel="noopener noreferrer">Investopedia: Personal Finance</a></li>
              <li><a href="https://www.consumerfinance.gov/" target="_blank" rel="noopener noreferrer">Consumer Financial Protection Bureau</a></li>
              <li><a href="https://www.mymoney.gov/" target="_blank" rel="noopener noreferrer">MyMoney.gov</a></li>
            </ul>
          </Typography>
        </Paper>
      </Box>
    </Container>
  );
}
