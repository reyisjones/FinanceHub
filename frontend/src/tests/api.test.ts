import { describe, it, expect, vi } from 'vitest';
import axios from 'axios';
import { getStockQuote, getCryptoPrices, getTopics } from '../services/api';

// Mock axios
vi.mock('axios');
const mockedAxios = axios as jest.Mocked<typeof axios>;

describe('API Service', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  describe('getTopics', () => {
    it('fetches finance topics successfully', async () => {
      const mockTopics = [
        { id: '1', title: 'Investments', description: 'Investment basics' },
        { id: '2', title: 'Stock Markets', description: 'Stock market overview' },
      ];

      mockedAxios.get.mockResolvedValueOnce({ data: mockTopics });

      const result = await getTopics();
      
      expect(result).toEqual(mockTopics);
      expect(mockedAxios.get).toHaveBeenCalledWith(
        expect.stringContaining('/topics')
      );
    });

    it('handles errors when fetching topics', async () => {
      mockedAxios.get.mockRejectedValueOnce(new Error('Network error'));

      await expect(getTopics()).rejects.toThrow('Network error');
    });
  });

  describe('getStockQuote', () => {
    it('fetches stock quote successfully', async () => {
      const mockQuote = {
        symbol: 'AAPL',
        price: 150.25,
        change: 2.5,
        changePercent: 1.69,
      };

      mockedAxios.get.mockResolvedValueOnce({ data: mockQuote });

      const result = await getStockQuote('AAPL');
      
      expect(result).toEqual(mockQuote);
      expect(mockedAxios.get).toHaveBeenCalledWith(
        expect.stringContaining('AAPL')
      );
    });

    it('validates symbol parameter', async () => {
      await expect(getStockQuote('')).rejects.toThrow();
    });
  });

  describe('getCryptoPrices', () => {
    it('fetches cryptocurrency prices successfully', async () => {
      const mockCrypto = [
        { id: 'bitcoin', symbol: 'BTC', price: 45000, change24h: 2.5 },
        { id: 'ethereum', symbol: 'ETH', price: 3000, change24h: 1.8 },
      ];

      mockedAxios.get.mockResolvedValueOnce({ data: mockCrypto });

      const result = await getCryptoPrices();
      
      expect(result).toEqual(mockCrypto);
      expect(result).toHaveLength(2);
    });

    it('handles empty response', async () => {
      mockedAxios.get.mockResolvedValueOnce({ data: [] });

      const result = await getCryptoPrices();
      
      expect(result).toEqual([]);
    });
  });
});
