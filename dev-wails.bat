@echo off
REM FinanceHub Wails Development Script for Windows
REM This script starts the Wails development server with hot reload

echo ======================================
echo FinanceHub - Wails Dev Server
echo ======================================
echo.

REM Check if Wails is installed
where wails >nul 2>nul
if %errorlevel% neq 0 (
    echo ERROR: Wails CLI is not installed!
    echo.
    echo Please install Wails first:
    echo   go install github.com/wailsapp/wails/v2/cmd/wails@latest
    echo.
    pause
    exit /b 1
)

echo Starting Wails development server...
echo Hot reload is enabled for frontend and backend changes
echo.
echo Press Ctrl+C to stop the server
echo.

wails dev

if %errorlevel% neq 0 (
    echo.
    echo ERROR: Failed to start development server!
    pause
    exit /b 1
)
