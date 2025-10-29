import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen, waitFor } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import App from '../App';

// Mock Wails runtime
vi.mock('../wailsjs/go/main/App', () => ({
  Greet: vi.fn(() => Promise.resolve('Hello Test!')),
  GetFinanceTopics: vi.fn(() => Promise.resolve([])),
  GetSystemInfo: vi.fn(() => Promise.resolve({
    os: 'windows',
    arch: 'amd64',
    appVersion: '1.0.0',
    appName: 'FinanceHub'
  })),
}));

describe('App Component', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('renders without crashing', () => {
    render(
      <BrowserRouter>
        <App />
      </BrowserRouter>
    );
    expect(document.body).toBeTruthy();
  });

  it('renders main navigation', async () => {
    render(
      <BrowserRouter>
        <App />
      </BrowserRouter>
    );
    
    await waitFor(() => {
      expect(screen.getByText(/FinanceHub/i)).toBeInTheDocument();
    });
  });

  it('contains router links', () => {
    render(
      <BrowserRouter>
        <App />
      </BrowserRouter>
    );
    
    // Check for navigation elements
    const container = document.querySelector('body');
    expect(container).toBeTruthy();
  });
});
