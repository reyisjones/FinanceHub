import { Box, Container, Typography, Grid, Card, CardContent, CardActionArea } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import TrendingUpIcon from '@mui/icons-material/TrendingUp';
import ShowChartIcon from '@mui/icons-material/ShowChart';
import AccountBalanceIcon from '@mui/icons-material/AccountBalance';
import CurrencyBitcoinIcon from '@mui/icons-material/CurrencyBitcoin';
import SavingsIcon from '@mui/icons-material/Savings';
import HomeWorkIcon from '@mui/icons-material/HomeWork';
import BusinessIcon from '@mui/icons-material/Business';
import PhoneIphoneIcon from '@mui/icons-material/PhoneIphone';
import SecurityIcon from '@mui/icons-material/Security';
import PublicIcon from '@mui/icons-material/Public';

const topics = [
  { title: 'Investments', path: '/investments', icon: TrendingUpIcon, description: 'Learn about various investment vehicles and strategies' },
  { title: 'Stock Markets', path: '/stock-markets', icon: ShowChartIcon, description: 'Understanding stock exchanges and equity trading' },
  { title: 'Bonds', path: '/bonds', icon: AccountBalanceIcon, description: 'Fixed-income securities and debt instruments' },
  { title: 'Cryptocurrencies', path: '/cryptocurrencies', icon: CurrencyBitcoinIcon, description: 'Digital currencies and blockchain technology' },
  { title: 'Personal Finance', path: '/personal-finance', icon: SavingsIcon, description: 'Budgeting, saving, and managing personal money' },
  { title: 'Real Estate', path: '/real-estate', icon: HomeWorkIcon, description: 'Property investment and real estate markets' },
  { title: 'Banking', path: '/banking', icon: BusinessIcon, description: 'Banking systems, accounts, and financial services' },
  { title: 'Fintech', path: '/fintech', icon: PhoneIphoneIcon, description: 'Financial technology and digital innovation' },
  { title: 'Risk Management', path: '/risk-management', icon: SecurityIcon, description: 'Identifying and mitigating financial risks' },
  { title: 'Global Economy', path: '/global-economy', icon: PublicIcon, description: 'International markets and macroeconomic trends' },
];

export default function HomePage() {
  const navigate = useNavigate();

  return (
    <Container maxWidth="lg">
      <Box sx={{ my: 4 }}>
        <Typography variant="h2" component="h1" gutterBottom align="center" sx={{ fontWeight: 700, mb: 2 }}>
          Welcome to FinanceHub
        </Typography>
        <Typography variant="h5" align="center" color="text.secondary" paragraph sx={{ mb: 6 }}>
          Your comprehensive platform for financial education and real-time market insights
        </Typography>

        <Box sx={{ mb: 6, p: 4, bgcolor: 'background.paper', borderRadius: 2, boxShadow: 2 }}>
          <Typography variant="h4" gutterBottom sx={{ fontWeight: 600 }}>
            About FinanceHub
          </Typography>
          <Typography variant="body1" paragraph>
            FinanceHub is designed to help you understand the world of finance through clear explanations and 
            live data from global financial markets. Whether you're a beginner learning the basics or an 
            experienced investor staying updated with market trends, we provide the tools and knowledge you need.
          </Typography>
          <Typography variant="body1" paragraph>
            Explore 10 key finance topics, each featuring:
          </Typography>
          <Typography component="div" variant="body1">
            <ul>
              <li>Educational summaries explaining core concepts</li>
              <li>Real-time data widgets showing live market information</li>
              <li>Interactive charts and visualizations</li>
              <li>Curated resources for deeper learning</li>
            </ul>
          </Typography>
        </Box>

        <Typography variant="h4" gutterBottom sx={{ fontWeight: 600, mb: 3 }}>
          Explore Finance Topics
        </Typography>

        <Grid container spacing={3}>
          {topics.map((topic) => {
            const IconComponent = topic.icon;
            return (
              <Grid item xs={12} sm={6} md={4} key={topic.path}>
                <Card sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
                  <CardActionArea 
                    onClick={() => navigate(topic.path)}
                    sx={{ flexGrow: 1, display: 'flex', flexDirection: 'column', alignItems: 'stretch' }}
                  >
                    <CardContent sx={{ textAlign: 'center', flexGrow: 1 }}>
                      <IconComponent sx={{ fontSize: 60, mb: 2, color: 'primary.main' }} />
                      <Typography gutterBottom variant="h5" component="h2" sx={{ fontWeight: 600 }}>
                        {topic.title}
                      </Typography>
                      <Typography variant="body2" color="text.secondary">
                        {topic.description}
                      </Typography>
                    </CardContent>
                  </CardActionArea>
                </Card>
              </Grid>
            );
          })}
        </Grid>
      </Box>
    </Container>
  );
}
