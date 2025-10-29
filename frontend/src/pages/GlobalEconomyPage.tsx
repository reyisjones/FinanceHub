import { Box, Container, Typography, Paper } from '@mui/material';

export default function GlobalEconomyPage() {
  return (
    <Container maxWidth="lg">
      <Box sx={{ my: 4 }}>
        <Typography variant="h3" component="h1" gutterBottom sx={{ fontWeight: 700 }}>
          Global Economy
        </Typography>
        <Paper sx={{ p: 3, mb: 4 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            About Global Economy
          </Typography>
          <Typography variant="body1" paragraph>
            The global economy encompasses worldwide economic activities, trade, and financial flows. 
            Key factors include GDP growth, inflation, unemployment, interest rates, currency exchange, and geopolitical events. 
            Understanding macroeconomic indicators and international markets helps investors make informed decisions and anticipate market movements.
          </Typography>
        </Paper>
        <Paper sx={{ p: 3 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            Learning Resources
          </Typography>
          <Typography component="div">
            <ul>
              <li><a href="https://www.investopedia.com/terms/g/globaleconomy.asp" target="_blank" rel="noopener noreferrer">Investopedia: Global Economy</a></li>
              <li><a href="https://www.imf.org/" target="_blank" rel="noopener noreferrer">International Monetary Fund</a></li>
              <li><a href="https://www.worldbank.org/" target="_blank" rel="noopener noreferrer">World Bank</a></li>
            </ul>
          </Typography>
        </Paper>
      </Box>
    </Container>
  );
}
