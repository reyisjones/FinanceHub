package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTopicsService(t *testing.T) {
	service := NewTopicsService()
	assert.NotNil(t, service, "TopicsService should not be nil")
}

func TestGetAllTopics(t *testing.T) {
	service := NewTopicsService()
	topics := service.GetAllTopics()

	assert.NotNil(t, topics, "Topics should not be nil")
	assert.Equal(t, 10, len(topics), "Should have exactly 10 topics")

	// Verify all topics have required fields
	for _, topic := range topics {
		assert.NotEmpty(t, topic.ID, "Topic ID should not be empty")
		assert.NotEmpty(t, topic.Title, "Topic title should not be empty")
		assert.NotEmpty(t, topic.Description, "Topic description should not be empty")
		assert.NotNil(t, topic.Keywords, "Topic keywords should not be nil")
	}
}

func TestGetTopicByID(t *testing.T) {
	service := NewTopicsService()

	tests := []struct {
		name        string
		topicID     string
		expectFound bool
	}{
		{
			name:        "Valid topic ID - investments",
			topicID:     "investments",
			expectFound: true,
		},
		{
			name:        "Valid topic ID - stock-markets",
			topicID:     "stock-markets",
			expectFound: true,
		},
		{
			name:        "Invalid topic ID",
			topicID:     "999",
			expectFound: false,
		},
		{
			name:        "Empty topic ID",
			topicID:     "",
			expectFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			topic := service.GetTopicByID(tt.topicID)

			if tt.expectFound {
				assert.NotNil(t, topic, "Topic should be found")
				assert.Equal(t, tt.topicID, topic.ID, "Topic ID should match")
			} else {
				assert.Nil(t, topic, "Topic should not be found")
			}
		})
	}
}

func TestTopicDataIntegrity(t *testing.T) {
	service := NewTopicsService()

	// Verify specific topics exist with expected titles
	expectedTopics := map[string]string{
		"investments":      "Investments",
		"stock-markets":    "Stock Markets",
		"bonds":            "Bonds",
		"cryptocurrencies": "Cryptocurrencies",
		"personal-finance": "Personal Finance",
		"real-estate":      "Real Estate",
		"banking":          "Banking",
		"fintech":          "Fintech",
		"risk-management":  "Risk Management",
		"global-economy":   "Global Economy",
	}

	for id, expectedTitle := range expectedTopics {
		topic := service.GetTopicByID(id)
		assert.NotNil(t, topic, "Topic %s should exist", id)
		if topic != nil {
			assert.Equal(t, expectedTitle, topic.Title, "Topic %s should have correct title", id)
		}
	}
}
