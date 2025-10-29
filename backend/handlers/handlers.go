package handlers

import (
	"net/http"

	"financehub/models"
	"financehub/services"

	"github.com/gin-gonic/gin"
)

// Handler holds all service dependencies
type Handler struct {
	AlphaVantage *services.AlphaVantageService
	CoinGecko    *services.CoinGeckoService
	Topics       *services.TopicsService
}

// NewHandler creates a new handler with all services
func NewHandler() *Handler {
	return &Handler{
		AlphaVantage: services.NewAlphaVantageService(),
		CoinGecko:    services.NewCoinGeckoService(),
		Topics:       services.NewTopicsService(),
	}
}

// GetAllTopics returns all finance topics
func (h *Handler) GetAllTopics(c *gin.Context) {
	topics := h.Topics.GetAllTopics()
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Data:    topics,
	})
}

// GetTopicByID returns a specific topic
func (h *Handler) GetTopicByID(c *gin.Context) {
	topicID := c.Param("id")
	topic := h.Topics.GetTopicByID(topicID)

	if topic == nil {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Success: false,
			Error:   "Topic not found",
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Data:    topic,
	})
}

// GetStockQuote returns stock quote for a symbol
func (h *Handler) GetStockQuote(c *gin.Context) {
	symbol := c.Param("symbol")

	quote, err := h.AlphaVantage.GetStockQuote(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Data:    quote,
	})
}

// GetStockTimeSeries returns historical stock data
func (h *Handler) GetStockTimeSeries(c *gin.Context) {
	symbol := c.Param("symbol")

	timeSeries, err := h.AlphaVantage.GetTimeSeriesDaily(symbol, 30)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Data:    timeSeries,
	})
}

// GetCryptoPrice returns cryptocurrency price
func (h *Handler) GetCryptoPrice(c *gin.Context) {
	coinID := c.Param("id")

	price, err := h.CoinGecko.GetCryptoPrice(coinID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Data:    price,
	})
}

// GetTopCryptos returns top cryptocurrencies
func (h *Handler) GetTopCryptos(c *gin.Context) {
	cryptos, err := h.CoinGecko.GetTopCryptos(10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Data:    cryptos,
	})
}

// GetCurrencyRate returns currency exchange rate
func (h *Handler) GetCurrencyRate(c *gin.Context) {
	from := c.Param("from")
	to := c.Param("to")

	rate, err := h.AlphaVantage.GetCurrencyExchangeRate(from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Data:    rate,
	})
}

// HealthCheck returns API health status
func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Finance Hub API is running",
	})
}
