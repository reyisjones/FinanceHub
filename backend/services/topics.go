package services

import (
	"financehub/models"
)

// TopicsService manages finance topic data
type TopicsService struct{}

// NewTopicsService creates a new topics service
func NewTopicsService() *TopicsService {
	return &TopicsService{}
}

// GetAllTopics returns all finance topics
func (s *TopicsService) GetAllTopics() []models.FinanceTopic {
	return []models.FinanceTopic{
		{
			ID:          "investments",
			Title:       "Investments",
			Description: "Learn about various investment vehicles and strategies",
			Summary: "Investment is the act of allocating resources, usually money, with the expectation of generating income or profit. " +
				"Understanding different investment types, risk levels, and strategies is crucial for building wealth over time. " +
				"Common investment vehicles include stocks, bonds, mutual funds, ETFs, and real estate.",
			Keywords: []string{"stocks", "bonds", "portfolio", "diversification", "ROI", "risk management"},
			Resources: []string{
				"https://www.investopedia.com/terms/i/investment.asp",
				"https://www.sec.gov/investor/pubs/investorpubs.htm",
				"https://www.investor.gov/introduction-investing",
			},
		},
		{
			ID:          "stock-markets",
			Title:       "Stock Markets",
			Description: "Understanding stock exchanges and equity trading",
			Summary: "Stock markets are venues where buyers and sellers meet to exchange equity shares of public corporations. " +
				"Major exchanges include NYSE, NASDAQ, and international markets. Stock prices fluctuate based on supply and demand, " +
				"company performance, economic indicators, and investor sentiment. Understanding market mechanics, indices, and trading strategies is essential for investors.",
			Keywords: []string{"NYSE", "NASDAQ", "trading", "equities", "indices", "bull market", "bear market"},
			Resources: []string{
				"https://www.investopedia.com/terms/s/stockmarket.asp",
				"https://www.sec.gov/fast-answers/answersmrktstruchtm.html",
				"https://finance.yahoo.com/",
			},
		},
		{
			ID:          "bonds",
			Title:       "Bonds",
			Description: "Fixed-income securities and debt instruments",
			Summary: "Bonds are debt securities where an investor loans money to an entity (government or corporation) for a defined period at a fixed interest rate. " +
				"They are considered safer than stocks but typically offer lower returns. Types include government bonds, corporate bonds, municipal bonds, and treasury securities. " +
				"Bond prices and yields move inversely, and factors like credit ratings, duration, and interest rates affect their value.",
			Keywords: []string{"fixed income", "yield", "treasury", "corporate bonds", "coupon rate", "maturity"},
			Resources: []string{
				"https://www.investopedia.com/terms/b/bond.asp",
				"https://www.investor.gov/introduction-investing/investing-basics/investment-products/bonds",
				"https://www.treasurydirect.gov/",
			},
		},
		{
			ID:          "cryptocurrencies",
			Title:       "Cryptocurrencies",
			Description: "Digital currencies and blockchain technology",
			Summary: "Cryptocurrencies are digital or virtual currencies that use cryptography for security and operate on decentralized networks based on blockchain technology. " +
				"Bitcoin, Ethereum, and thousands of altcoins exist. Crypto offers potential for high returns but comes with significant volatility and risk. " +
				"Key concepts include blockchain, mining, wallets, decentralized finance (DeFi), and smart contracts.",
			Keywords: []string{"Bitcoin", "Ethereum", "blockchain", "DeFi", "altcoins", "crypto wallet", "mining"},
			Resources: []string{
				"https://www.investopedia.com/terms/c/cryptocurrency.asp",
				"https://www.coingecko.com/",
				"https://ethereum.org/en/what-is-ethereum/",
			},
		},
		{
			ID:          "personal-finance",
			Title:       "Personal Finance",
			Description: "Budgeting, saving, and managing personal money",
			Summary: "Personal finance encompasses budgeting, saving, investing, insurance, mortgages, retirement planning, and tax planning. " +
				"Key principles include living below your means, building an emergency fund, eliminating high-interest debt, and saving for retirement. " +
				"Financial literacy helps individuals make informed decisions about spending, saving, and investing for long-term financial security.",
			Keywords: []string{"budgeting", "savings", "retirement planning", "emergency fund", "debt management", "financial literacy"},
			Resources: []string{
				"https://www.investopedia.com/personal-finance-4427760",
				"https://www.consumerfinance.gov/",
				"https://www.mymoney.gov/",
			},
		},
		{
			ID:          "real-estate",
			Title:       "Real Estate",
			Description: "Property investment and real estate markets",
			Summary: "Real estate involves buying, selling, and investing in property—residential, commercial, or industrial. " +
				"It's a tangible asset that can provide rental income, appreciation, and tax benefits. Real estate investment trusts (REITs) allow indirect investment. " +
				"Factors affecting real estate include location, market conditions, interest rates, and economic trends.",
			Keywords: []string{"property", "REIT", "mortgage", "rental income", "appreciation", "commercial real estate"},
			Resources: []string{
				"https://www.investopedia.com/terms/r/realestate.asp",
				"https://www.nar.realtor/",
				"https://www.zillow.com/research/",
			},
		},
		{
			ID:          "banking",
			Title:       "Banking",
			Description: "Banking systems, accounts, and financial services",
			Summary: "Banking institutions provide financial services including deposits, loans, credit, and payment processing. " +
				"Types of accounts include checking, savings, money market, and certificates of deposit (CDs). Banks are regulated to ensure stability and protect consumers. " +
				"Understanding banking products, fees, interest rates, and digital banking tools is important for managing finances effectively.",
			Keywords: []string{"checking account", "savings account", "loans", "credit", "interest rates", "online banking"},
			Resources: []string{
				"https://www.investopedia.com/terms/b/bank.asp",
				"https://www.fdic.gov/",
				"https://www.bankrate.com/",
			},
		},
		{
			ID:          "fintech",
			Title:       "Fintech",
			Description: "Financial technology and digital innovation",
			Summary: "Fintech (financial technology) refers to innovations that aim to compete with or enhance traditional financial services. " +
				"This includes mobile banking, peer-to-peer payments, robo-advisors, blockchain, and digital currencies. " +
				"Fintech improves accessibility, reduces costs, and increases efficiency in financial services through technology.",
			Keywords: []string{"mobile payments", "robo-advisors", "digital wallets", "peer-to-peer lending", "insurtech", "neobanks"},
			Resources: []string{
				"https://www.investopedia.com/terms/f/fintech.asp",
				"https://www.mckinsey.com/industries/financial-services/our-insights",
				"https://techcrunch.com/tag/fintech/",
			},
		},
		{
			ID:          "risk-management",
			Title:       "Risk Management",
			Description: "Identifying and mitigating financial risks",
			Summary: "Risk management in finance involves identifying, analyzing, and mitigating potential losses in investments and financial decisions. " +
				"Strategies include diversification, hedging, insurance, and asset allocation. Understanding risk tolerance, risk-return tradeoff, and various types of risk " +
				"(market risk, credit risk, liquidity risk) is essential for protecting and growing wealth.",
			Keywords: []string{"diversification", "hedging", "insurance", "risk assessment", "volatility", "asset allocation"},
			Resources: []string{
				"https://www.investopedia.com/terms/r/riskmanagement.asp",
				"https://www.cfa.org/",
				"https://www.prmia.org/",
			},
		},
		{
			ID:          "global-economy",
			Title:       "Global Economy",
			Description: "International markets and macroeconomic trends",
			Summary: "The global economy encompasses worldwide economic activities, trade, and financial flows. " +
				"Key factors include GDP growth, inflation, unemployment, interest rates, currency exchange, and geopolitical events. " +
				"Understanding macroeconomic indicators and international markets helps investors make informed decisions and anticipate market movements.",
			Keywords: []string{"GDP", "inflation", "monetary policy", "trade", "emerging markets", "forex", "central banks"},
			Resources: []string{
				"https://www.investopedia.com/terms/g/globaleconomy.asp",
				"https://www.imf.org/",
				"https://www.worldbank.org/",
			},
		},
	}
}

// GetTopicByID returns a specific topic by ID
func (s *TopicsService) GetTopicByID(id string) *models.FinanceTopic {
	topics := s.GetAllTopics()
	for _, topic := range topics {
		if topic.ID == id {
			return &topic
		}
	}
	return nil
}
