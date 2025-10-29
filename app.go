package main

import (
	"context"
	"financehub/models"
	"financehub/services"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// App struct
type App struct {
	ctx           context.Context
	topicsService *services.TopicsService
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		topicsService: services.NewTopicsService(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a personalized greeting for the user
func (a *App) Greet(name string) string {
	if name == "" {
		name = "Guest"
	}
	return fmt.Sprintf("Hello %s! Welcome to FinanceHub - Your gateway to financial knowledge and market insights.", name)
}

// GetFinanceTopics returns all 10 finance topics available in the application
func (a *App) GetFinanceTopics() []models.FinanceTopic {
	return a.topicsService.GetAllTopics()
}

// GetTopicByID returns a specific finance topic by its ID
func (a *App) GetTopicByID(id string) *models.FinanceTopic {
	return a.topicsService.GetTopicByID(id)
}

// GetSystemInfo returns information about the system running the app
func (a *App) GetSystemInfo() map[string]string {
	version := a.GetAppVersion()
	return map[string]string{
		"os":         runtime.GOOS,
		"arch":       runtime.GOARCH,
		"goVersion":  runtime.Version(),
		"numCPU":     fmt.Sprintf("%d", runtime.NumCPU()),
		"appVersion": version,
		"appName":    "FinanceHub",
	}
}

// GetAppVersion returns the current application version
func (a *App) GetAppVersion() string {
	// Try to read from VERSION file
	if content, err := os.ReadFile("VERSION"); err == nil {
		return strings.TrimSpace(string(content))
	}
	// Fallback version
	return "1.0.0"
}

// IsProduction returns whether the app is running in production mode
func (a *App) IsProduction() bool {
	// This can be set via build tags or environment variables
	return true
}
