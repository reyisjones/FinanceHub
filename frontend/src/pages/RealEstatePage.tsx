import { Box, Container, Typography, Paper } from '@mui/material';

export default function RealEstatePage() {
  return (
    <Container maxWidth="lg">
      <Box sx={{ my: 4 }}>
        <Typography variant="h3" component="h1" gutterBottom sx={{ fontWeight: 700 }}>
          Real Estate
        </Typography>
        <Paper sx={{ p: 3, mb: 4 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            About Real Estate
          </Typography>
          <Typography variant="body1" paragraph>
            Real estate involves buying, selling, and investing in property—residential, commercial, or industrial. 
            It's a tangible asset that can provide rental income, appreciation, and tax benefits. Real estate investment trusts (REITs) allow indirect investment. 
            Factors affecting real estate include location, market conditions, interest rates, and economic trends.
          </Typography>
        </Paper>
        <Paper sx={{ p: 3 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            Learning Resources
          </Typography>
          <Typography component="div">
            <ul>
              <li><a href="https://www.investopedia.com/terms/r/realestate.asp" target="_blank" rel="noopener noreferrer">Investopedia: Real Estate</a></li>
              <li><a href="https://www.nar.realtor/" target="_blank" rel="noopener noreferrer">National Association of Realtors</a></li>
              <li><a href="https://www.zillow.com/research/" target="_blank" rel="noopener noreferrer">Zillow Research</a></li>
            </ul>
          </Typography>
        </Paper>
      </Box>
    </Container>
  );
}
