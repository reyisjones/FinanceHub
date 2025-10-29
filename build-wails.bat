@echo off
REM FinanceHub Wails Build Script for Windows
REM This script builds the FinanceHub desktop application using Wails

echo ======================================
echo FinanceHub - Wails Build Script
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

echo [1/5] Checking dependencies...
echo.

REM Check Go installation
where go >nul 2>nul
if %errorlevel% neq 0 (
    echo ERROR: Go is not installed!
    echo Please install Go 1.21+ from https://golang.org/dl/
    pause
    exit /b 1
)

REM Check Node.js installation
where node >nul 2>nul
if %errorlevel% neq 0 (
    echo ERROR: Node.js is not installed!
    echo Please install Node.js 18+ from https://nodejs.org/
    pause
    exit /b 1
)

echo [2/5] Installing frontend dependencies...
cd frontend
call npm install
if %errorlevel% neq 0 (
    echo ERROR: Failed to install frontend dependencies!
    pause
    exit /b 1
)
cd ..

echo.
echo [3/5] Downloading Go modules...
go mod download
if %errorlevel% neq 0 (
    echo ERROR: Failed to download Go modules!
    pause
    exit /b 1
)

echo.
echo [4/5] Building FinanceHub application...
echo This may take several minutes...
echo.

wails build -clean
if %errorlevel% neq 0 (
    echo ERROR: Build failed!
    pause
    exit /b 1
)

echo.
echo [5/5] Build completed successfully!
echo.
echo ======================================
echo Build artifacts location:
echo   build\bin\FinanceHub.exe
echo ======================================
echo.
echo You can now run the application from build\bin\FinanceHub.exe
echo.
pause
