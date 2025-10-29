import { Box, Container, Typography, Paper } from '@mui/material';

export default function FintechPage() {
  return (
    <Container maxWidth="lg">
      <Box sx={{ my: 4 }}>
        <Typography variant="h3" component="h1" gutterBottom sx={{ fontWeight: 700 }}>
          Fintech
        </Typography>
        <Paper sx={{ p: 3, mb: 4 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            About Fintech
          </Typography>
          <Typography variant="body1" paragraph>
            Fintech (financial technology) refers to innovations that aim to compete with or enhance traditional financial services. 
            This includes mobile banking, peer-to-peer payments, robo-advisors, blockchain, and digital currencies. 
            Fintech improves accessibility, reduces costs, and increases efficiency in financial services through technology.
          </Typography>
        </Paper>
        <Paper sx={{ p: 3 }}>
          <Typography variant="h5" gutterBottom sx={{ fontWeight: 600 }}>
            Learning Resources
          </Typography>
          <Typography component="div">
            <ul>
              <li><a href="https://www.investopedia.com/terms/f/fintech.asp" target="_blank" rel="noopener noreferrer">Investopedia: Fintech</a></li>
              <li><a href="https://www.mckinsey.com/industries/financial-services/our-insights" target="_blank" rel="noopener noreferrer">McKinsey Financial Services</a></li>
              <li><a href="https://techcrunch.com/tag/fintech/" target="_blank" rel="noopener noreferrer">TechCrunch Fintech</a></li>
            </ul>
          </Typography>
        </Paper>
      </Box>
    </Container>
  );
}
