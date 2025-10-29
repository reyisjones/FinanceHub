@echo off
echo ======================================
echo Starting FinanceHub Development Servers
echo ======================================
echo.

REM Start Backend in new window
echo Starting Backend server...
start "FinanceHub Backend" cmd /k "cd backend && go run main.go"

REM Wait a bit for backend to start
timeout /t 3 /nobreak >nul

REM Start Frontend in new window
echo Starting Frontend server...
start "FinanceHub Frontend" cmd /k "cd frontend && npm run dev"

echo.
echo ======================================
echo Both servers are starting!
echo ======================================
echo.
echo Backend: http://localhost:8080
echo Frontend: http://localhost:5173
echo.
echo Press any key to stop all servers...
pause >nul

REM Kill both processes when user presses a key
taskkill /FI "WINDOWTITLE eq FinanceHub Backend" /F >nul 2>&1
taskkill /FI "WINDOWTITLE eq FinanceHub Frontend" /F >nul 2>&1

echo Servers stopped.
