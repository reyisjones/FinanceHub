import axios from 'axios';
import type { APIResponse, FinanceTopic, StockQuote, CryptoPrice, CurrencyRate, TimeSeriesData } from '../types';

const API_BASE_URL = 'http://localhost:8080/api';

const api = axios.create({
  baseURL: API_BASE_URL,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const financeAPI = {
  // Topics
  getAllTopics: async (): Promise<FinanceTopic[]> => {
    const response = await api.get<APIResponse<FinanceTopic[]>>('/topics');
    return response.data.data || [];
  },

  getTopicById: async (id: string): Promise<FinanceTopic> => {
    const response = await api.get<APIResponse<FinanceTopic>>(`/topics/${id}`);
    if (!response.data.data) {
      throw new Error('Topic not found');
    }
    return response.data.data;
  },

  // Stocks
  getStockQuote: async (symbol: string): Promise<StockQuote> => {
    const response = await api.get<APIResponse<StockQuote>>(`/stocks/${symbol}`);
    if (!response.data.data) {
      throw new Error('Stock data not found');
    }
    return response.data.data;
  },

  getStockTimeSeries: async (symbol: string): Promise<TimeSeriesData[]> => {
    const response = await api.get<APIResponse<TimeSeriesData[]>>(`/stocks/${symbol}/timeseries`);
    return response.data.data || [];
  },

  // Cryptocurrencies
  getTopCryptos: async (): Promise<CryptoPrice[]> => {
    const response = await api.get<APIResponse<CryptoPrice[]>>('/crypto/top');
    return response.data.data || [];
  },

  getCryptoPrice: async (id: string): Promise<CryptoPrice> => {
    const response = await api.get<APIResponse<CryptoPrice>>(`/crypto/${id}`);
    if (!response.data.data) {
      throw new Error('Crypto data not found');
    }
    return response.data.data;
  },

  // Currency
  getCurrencyRate: async (from: string, to: string): Promise<CurrencyRate> => {
    const response = await api.get<APIResponse<CurrencyRate>>(`/currency/${from}/${to}`);
    if (!response.data.data) {
      throw new Error('Currency rate not found');
    }
    return response.data.data;
  },
};

export default financeAPI;
