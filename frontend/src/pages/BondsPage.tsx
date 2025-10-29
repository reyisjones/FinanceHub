import { Box, Container, Typography, Paper } from '@mui/material';

export default function BondsPage() {
  return (
    <Container maxWidth="lg">
      <Box sx={{ my: 4 }}>
        <Typography variant="h3" component="h1" gutterBottom sx={{ fontWeight: 700 }}>
          Bonds
        </Typography>
        <Paper sx={{ p: 3, mb: 4 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            About Bonds
          </Typography>
          <Typography variant="body1" paragraph>
            Bonds are debt securities where an investor loans money to an entity (government or corporation) for a defined period at a fixed interest rate. 
            They are considered safer than stocks but typically offer lower returns. Types include government bonds, corporate bonds, municipal bonds, and treasury securities.
          </Typography>
        </Paper>
        <Paper sx={{ p: 3 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            Learning Resources
          </Typography>
          <Typography component="div">
            <ul>
              <li><a href="https://www.investopedia.com/terms/b/bond.asp" target="_blank" rel="noopener noreferrer">Investopedia: Bonds</a></li>
              <li><a href="https://www.investor.gov/introduction-investing/investing-basics/investment-products/bonds" target="_blank" rel="noopener noreferrer">Investor.gov: Bonds</a></li>
              <li><a href="https://www.treasurydirect.gov/" target="_blank" rel="noopener noreferrer">TreasuryDirect</a></li>
            </ul>
          </Typography>
        </Paper>
      </Box>
    </Container>
  );
}
