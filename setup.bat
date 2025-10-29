@echo off
echo ======================================
echo FinanceHub - Initial Setup
echo ======================================
echo.

REM Check if Go is installed
where go >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Go is not installed!
    echo Please install Go from https://golang.org/dl/
    pause
    exit /b 1
)

REM Check if Node.js is installed
where node >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Node.js is not installed!
    echo Please install Node.js from https://nodejs.org/
    pause
    exit /b 1
)

echo [OK] Go and Node.js are installed
echo.

REM Setup Backend
echo ======================================
echo Setting up Backend...
echo ======================================
cd backend

if not exist .env (
    echo Creating .env file from template...
    copy .env.example .env
    echo.
    echo [IMPORTANT] Please edit backend\.env and add your Alpha Vantage API key
    echo Get your free API key at: https://www.alphavantage.co/support/#api-key
    echo.
    pause
)

echo Installing Go dependencies...
go mod download
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Failed to install Go dependencies
    pause
    exit /b 1
)

echo [OK] Backend setup complete!
echo.

REM Setup Frontend
echo ======================================
echo Setting up Frontend...
echo ======================================
cd ..\frontend

echo Installing npm dependencies (this may take a few minutes)...
call npm install
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Failed to install npm dependencies
    pause
    exit /b 1
)

echo [OK] Frontend setup complete!
echo.

cd ..

echo ======================================
echo Setup Complete!
echo ======================================
echo.
echo Next steps:
echo 1. Make sure you added your Alpha Vantage API key to backend\.env
echo 2. Run 'start-dev.bat' to start both servers
echo.
pause
