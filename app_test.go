package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	app := NewApp()
	assert.NotNil(t, app, "App should not be nil")
	assert.NotNil(t, app.topicsService, "TopicsService should be initialized")
}

func TestAppStartup(t *testing.T) {
	app := NewApp()
	ctx := context.Background()

	app.startup(ctx)

	assert.NotNil(t, app.ctx, "Context should be set after startup")
}

func TestGreet(t *testing.T) {
	app := NewApp()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Greet with name",
			input:    "John",
			expected: "Hello John! Welcome to FinanceHub!",
		},
		{
			name:     "Greet with empty name",
			input:    "",
			expected: "Hello ! Welcome to FinanceHub!",
		},
		{
			name:     "Greet with special characters",
			input:    "José",
			expected: "Hello José! Welcome to FinanceHub!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := app.Greet(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetFinanceTopics(t *testing.T) {
	app := NewApp()

	topics := app.GetFinanceTopics()

	assert.NotNil(t, topics, "Topics should not be nil")
	assert.Equal(t, 10, len(topics), "Should return 10 topics")

	// Verify first topic structure
	if len(topics) > 0 {
		assert.NotEmpty(t, topics[0].ID)
		assert.NotEmpty(t, topics[0].Title)
		assert.NotEmpty(t, topics[0].Description)
	}
}

func TestGetTopicByID(t *testing.T) {
	app := NewApp()

	tests := []struct {
		name       string
		topicID    string
		shouldFind bool
	}{
		{
			name:       "Valid topic ID",
			topicID:    "1",
			shouldFind: true,
		},
		{
			name:       "Another valid ID",
			topicID:    "5",
			shouldFind: true,
		},
		{
			name:       "Invalid topic ID",
			topicID:    "999",
			shouldFind: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			topic := app.GetTopicByID(tt.topicID)

			if tt.shouldFind {
				assert.NotNil(t, topic, "Topic should be found")
				assert.Equal(t, tt.topicID, topic.ID)
			} else {
				assert.Nil(t, topic, "Topic should not be found")
			}
		})
	}
}

func TestGetSystemInfo(t *testing.T) {
	app := NewApp()

	info := app.GetSystemInfo()

	assert.NotNil(t, info, "System info should not be nil")
	assert.NotEmpty(t, info["os"], "OS should be set")
	assert.NotEmpty(t, info["arch"], "Architecture should be set")
	assert.NotEmpty(t, info["appVersion"], "App version should be set")
	assert.Equal(t, "FinanceHub", info["appName"])
}

func TestGetAppVersion(t *testing.T) {
	app := NewApp()

	version := app.GetAppVersion()

	assert.NotEmpty(t, version, "Version should not be empty")
	assert.Regexp(t, `^\d+\.\d+\.\d+$`, version, "Version should follow semver format")
}

func TestIsProduction(t *testing.T) {
	app := NewApp()

	isProd := app.IsProduction()

	assert.IsType(t, true, isProd, "IsProduction should return a boolean")
}
