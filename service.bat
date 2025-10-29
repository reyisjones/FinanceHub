@echo off
REM FinanceHub Windows Service Management Script
REM This script manages the FinanceHub Windows Service

echo ======================================
echo FinanceHub - Service Manager
echo ======================================
echo.

if "%1"=="" (
    echo Usage: service.bat [install^|uninstall^|start^|stop^|restart]
    echo.
    echo Commands:
    echo   install   - Install FinanceHub as a Windows Service
    echo   uninstall - Remove FinanceHub Windows Service
    echo   start     - Start the FinanceHub service
    echo   stop      - Stop the FinanceHub service
    echo   restart   - Restart the FinanceHub service
    echo.
    pause
    exit /b 1
)

REM Check if running as administrator
net session >nul 2>&1
if %errorlevel% neq 0 (
    echo ERROR: This script requires Administrator privileges!
    echo Please run as Administrator.
    echo.
    pause
    exit /b 1
)

set SERVICE_EXE=build\bin\FinanceHub.exe

if not exist "%SERVICE_EXE%" (
    echo ERROR: FinanceHub.exe not found in build\bin\
    echo Please build the application first using build-wails.bat
    echo.
    pause
    exit /b 1
)

if "%1"=="install" (
    echo Installing FinanceHub as Windows Service...
    "%SERVICE_EXE%" install
    if %errorlevel% equ 0 (
        echo.
        echo ✓ Service installed successfully!
        echo   Service will start automatically with Windows
    ) else (
        echo.
        echo ✗ Service installation failed!
    )
)

if "%1"=="uninstall" (
    echo Uninstalling FinanceHub Windows Service...
    "%SERVICE_EXE%" uninstall
    if %errorlevel% equ 0 (
        echo.
        echo ✓ Service uninstalled successfully!
    ) else (
        echo.
        echo ✗ Service uninstallation failed!
    )
)

if "%1"=="start" (
    echo Starting FinanceHub service...
    "%SERVICE_EXE%" start
    if %errorlevel% equ 0 (
        echo.
        echo ✓ Service started successfully!
    ) else (
        echo.
        echo ✗ Failed to start service!
    )
)

if "%1"=="stop" (
    echo Stopping FinanceHub service...
    "%SERVICE_EXE%" stop
    if %errorlevel% equ 0 (
        echo.
        echo ✓ Service stopped successfully!
    ) else (
        echo.
        echo ✗ Failed to stop service!
    )
)

if "%1"=="restart" (
    echo Restarting FinanceHub service...
    "%SERVICE_EXE%" restart
    if %errorlevel% equ 0 (
        echo.
        echo ✓ Service restarted successfully!
    ) else (
        echo.
        echo ✗ Failed to restart service!
    )
)

echo.
pause
