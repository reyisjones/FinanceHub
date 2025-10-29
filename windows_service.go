package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/kardianos/service"
)

// Program structures.
type program struct {
	logger service.Logger
}

// Start initializes and starts the FinanceHub service
func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

// run executes the main service logic
func (p *program) run() {
	p.logger.Info("FinanceHub service is running")

	// Log service start event
	logServiceEvent("Service started successfully")

	// Here you would start your Wails application or any background tasks
	// For a true background service, you might want to run specific tasks
	// rather than the full GUI application

	// Keep the service running
	for {
		// Perform periodic health checks or maintenance tasks
		time.Sleep(30 * time.Second)
		p.logger.Info("FinanceHub service heartbeat")
	}
}

// Stop terminates the service gracefully
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return within a few seconds.
	p.logger.Info("FinanceHub service is stopping")
	logServiceEvent("Service stopped gracefully")
	return nil
}

// logServiceEvent writes service lifecycle events to a log file
func logServiceEvent(message string) {
	// Get executable directory for log file location
	exePath, err := os.Executable()
	if err != nil {
		return
	}

	logDir := filepath.Join(filepath.Dir(exePath), "logs")
	os.MkdirAll(logDir, 0755)

	logFile := filepath.Join(logDir, "financehub_service.log")
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] %s\n", timestamp, message)
	f.WriteString(logEntry)
}

// RunAsService configures and runs FinanceHub as a Windows service
func RunAsService() {
	svcConfig := &service.Config{
		Name:        "FinanceHubService",
		DisplayName: "FinanceHub Service",
		Description: "FinanceHub - Finance Learning & Insights Desktop Application Service. Provides background data synchronization and system integration.",
		Option: service.KeyValue{
			"StartType": "automatic",
		},
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	logger, err := s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	prg.logger = logger

	// Handle service control commands
	if len(os.Args) > 1 {
		cmd := os.Args[1]
		switch cmd {
		case "install":
			err := s.Install()
			if err != nil {
				logger.Errorf("Failed to install service: %v", err)
				return
			}
			logger.Info("Service installed successfully")
			logServiceEvent("Service installed")
			return
		case "uninstall":
			err := s.Uninstall()
			if err != nil {
				logger.Errorf("Failed to uninstall service: %v", err)
				return
			}
			logger.Info("Service uninstalled successfully")
			logServiceEvent("Service uninstalled")
			return
		case "start":
			err := s.Start()
			if err != nil {
				logger.Errorf("Failed to start service: %v", err)
				return
			}
			logger.Info("Service started successfully")
			return
		case "stop":
			err := s.Stop()
			if err != nil {
				logger.Errorf("Failed to stop service: %v", err)
				return
			}
			logger.Info("Service stopped successfully")
			return
		case "restart":
			err := s.Restart()
			if err != nil {
				logger.Errorf("Failed to restart service: %v", err)
				return
			}
			logger.Info("Service restarted successfully")
			return
		}
	}

	// Run the service
	err = s.Run()
	if err != nil {
		logger.Error(err)
		logServiceEvent(fmt.Sprintf("Service error: %v", err))
	}
}
