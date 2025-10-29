export interface FinanceTopic {
  id: string;
  title: string;
  description: string;
  summary: string;
  keywords: string[];
  resources: string[];
}

export interface StockQuote {
  symbol: string;
  price: number;
  change: number;
  changePercent: number;
  volume: number;
  high: number;
  low: number;
  open: number;
  previousClose: number;
  lastUpdated: string;
}

export interface CryptoPrice {
  id: string;
  symbol: string;
  name: string;
  currentPrice: number;
  marketCap: number;
  marketCapRank: number;
  priceChange24h: number;
  priceChangePercent24h: number;
  high24h: number;
  low24h: number;
  circulatingSupply: number;
  lastUpdated: string;
}

export interface CurrencyRate {
  fromCurrency: string;
  toCurrency: string;
  rate: number;
  bid: number;
  ask: number;
  lastUpdated: string;
}

export interface TimeSeriesData {
  date: string;
  open: number;
  high: number;
  low: number;
  close: number;
  volume: number;
}

export interface APIResponse<T = any> {
  success: boolean;
  data?: T;
  error?: string;
  message?: string;
}
