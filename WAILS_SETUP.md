# FinanceHub - Wails Desktop Application Setup Guide

Complete guide for building, running, and deploying FinanceHub as a desktop application using Wails.

## 📋 Table of Contents

- [Prerequisites](#prerequisites)
- [Initial Setup](#initial-setup)
- [Development Workflow](#development-workflow)
- [Building for Production](#building-for-production)
- [Windows Service Installation](#windows-service-installation)
- [Code Signing](#code-signing)
- [GitHub Actions CI/CD](#github-actions-cicd)
- [Troubleshooting](#troubleshooting)

## 🔧 Prerequisites

### Required Software

1. **Go 1.21 or higher**
   - Download: https://golang.org/dl/
   - Verify: `go version`

2. **Node.js 18+ and npm**
   - Download: https://nodejs.org/
   - Verify: `node --version` and `npm --version`

3. **Wails CLI**
   ```bash
   go install github.com/wailsapp/wails/v2/cmd/wails@latest
   ```
   - Verify: `wails version`

4. **Alpha Vantage API Key**
   - Get free key: https://www.alphavantage.co/support/#api-key

### Platform-Specific Requirements

#### Windows
- Windows 10/11 (64-bit)
- WebView2 Runtime (usually pre-installed on Windows 11)
- Download: https://developer.microsoft.com/en-us/microsoft-edge/webview2/

#### Optional Tools
- **NSIS** (for creating installers)
  ```bash
  choco install nsis
  ```
- **Code Signing Certificate** (for production releases)

## 🚀 Initial Setup

### 1. Clone and Configure

```bash
# Clone the repository
git clone https://github.com/reyisjones/FinanceHub.git
cd FinanceHub

# Create .env file for API keys (if using backend separately)
cd backend
cp .env.example .env
# Edit .env and add your Alpha Vantage API key
```

### 2. Install Dependencies

**Option A: Using Wails CLI (Recommended)**
```bash
# From project root
wails build -clean
```
This automatically installs both Go and npm dependencies.

**Option B: Manual Installation**
```bash
# Install frontend dependencies
cd frontend
npm install
cd ..

# Download Go modules
go mod download
```

### 3. Generate Wails Bindings

The Wails CLI automatically generates TypeScript bindings for your Go methods:

```bash
wails dev
```

This creates the `frontend/wailsjs` directory with TypeScript definitions.

## 💻 Development Workflow

### Starting Development Server

**Option A: Using Wails Dev Server (Recommended)**
```bash
# Windows
dev-wails.bat

# Or directly
wails dev
```

**Option B: Separate Frontend and Backend**
```bash
# Terminal 1: Start backend API server
cd backend
go run main.go

# Terminal 2: Start frontend dev server
cd frontend
npm run dev
```

### Development Features

- **Hot Reload**: Frontend changes reload automatically
- **Backend Reload**: Go code changes restart the app
- **Dev Tools**: Press F12 to open browser developer tools
- **API Testing**: Access Wails test page at `/wails-test`

### Available Wails Bindings

The following Go methods are available in the frontend:

```typescript
import { Greet, GetFinanceTopics, GetTopicByID, GetSystemInfo, GetAppVersion } from '../wailsjs/go/main/App';

// Greet user
const message = await Greet("John");

// Get all finance topics
const topics = await GetFinanceTopics();

// Get specific topic
const topic = await GetTopicByID(1);

// Get system information
const sysInfo = await GetSystemInfo();

// Get app version
const version = await GetAppVersion();
```

## 🏗️ Building for Production

### Quick Build

```bash
# Windows
build-wails.bat

# Or directly
wails build -clean
```

The executable will be created at: `build/bin/FinanceHub.exe`

### Build Options

**Debug Build (with console)**
```bash
wails build -debug
```

**Production Build (optimized, no console)**
```bash
wails build -clean -production
```

**Platform-Specific Build**
```bash
# Windows AMD64
wails build -platform windows/amd64

# Windows ARM64
wails build -platform windows/arm64
```

### Build Configuration

Edit `wails.json` to customize build settings:

```json
{
  "name": "FinanceHub",
  "outputfilename": "FinanceHub",
  "frontend:install": "cd frontend && npm install",
  "frontend:build": "cd frontend && npm run build",
  "windows": {
    "hideConsole": true,
    "minWidth": 800,
    "minHeight": 600,
    "theme": "system"
  }
}
```

## 🔧 Windows Service Installation

FinanceHub can run as a Windows background service that starts automatically with Windows.

### Install Service

```bash
# Run as Administrator
service.bat install
```

### Manage Service

```bash
# Start service
service.bat start

# Stop service
service.bat stop

# Restart service
service.bat restart

# Uninstall service
service.bat uninstall
```

### Service Logs

Service logs are stored in:
```
<installation_directory>/logs/financehub_service.log
```

### Service Features

- ✅ Auto-start with Windows
- ✅ Runs in background (no console window)
- ✅ Lifecycle event logging
- ✅ Health monitoring and heartbeat
- ✅ Graceful shutdown handling

## 🔐 Code Signing

### Prerequisites

1. **Obtain a Code Signing Certificate**
   - Use a trusted CA (DigiCert, Sectigo, etc.)
   - WebTrust-certified for Windows
   - Standard or EV (Extended Validation) certificate

2. **Export Certificate**
   ```powershell
   # Export as PFX with private key
   # Password protect the certificate
   ```

### Manual Signing

```powershell
# Sign the executable
signtool sign /f "path\to\certificate.pfx" /p "password" /t http://timestamp.digicert.com "build\bin\FinanceHub.exe"

# Verify signature
signtool verify /pa "build\bin\FinanceHub.exe"
```

### GitHub Actions Signing

Set up repository secrets:

1. Go to Repository Settings → Secrets → Actions
2. Add secrets:
   - `WINDOWS_CERTIFICATE_BASE64`: Base64-encoded PFX file
   - `WINDOWS_CERTIFICATE_PASSWORD`: Certificate password

**Create Base64 Certificate**
```powershell
$bytes = [System.IO.File]::ReadAllBytes("certificate.pfx")
$base64 = [System.Convert]::ToBase64String($bytes)
Set-Content -Path "certificate_base64.txt" -Value $base64
```

The GitHub Actions workflow will automatically sign releases.

## 🚀 GitHub Actions CI/CD

### Automated Release Workflow

The `.github/workflows/build.yml` workflow automatically:

1. ✅ Builds the application for Windows
2. ✅ Signs the executable (if certificate is configured)
3. ✅ Creates installer with NSIS
4. ✅ Generates release notes from commits
5. ✅ Creates GitHub release with versioned artifacts
6. ✅ Uploads build artifacts

### Triggering a Release

**Option A: Create Git Tag**
```bash
git tag v1.0.0
git push origin v1.0.0
```

**Option B: Manual Workflow Dispatch**
1. Go to Actions tab in GitHub
2. Select "Build and Release FinanceHub"
3. Click "Run workflow"
4. Enter version number (e.g., 1.0.0)

### Release Artifacts

Each release includes:
- `FinanceHub.exe` - Signed executable
- `FinanceHub-Setup-v1.0.0.exe` - NSIS installer
- `FinanceHub-Windows-x64-v1.0.0.zip` - Portable version
- `FinanceHub-v1.0.0.sha256` - Checksum file

### Versioning

Follow semantic versioning:
- `v1.0.0` - Major release
- `v1.1.0` - Feature update
- `v1.0.1` - Bug fix/patch

## 🐛 Troubleshooting

### Wails Not Found
```bash
# Reinstall Wails CLI
go install github.com/wailsapp/wails/v2/cmd/wails@latest

# Add Go bin to PATH
# Windows: Add %USERPROFILE%\go\bin to PATH
```

### WebView2 Missing
```
Error: WebView2 runtime not found
```
**Solution**: Install WebView2 Runtime
- Download: https://developer.microsoft.com/en-us/microsoft-edge/webview2/

### Frontend Build Fails
```bash
# Clear npm cache
cd frontend
rm -rf node_modules package-lock.json
npm install
```

### Backend Build Fails
```bash
# Clear Go module cache
go clean -modcache
go mod download
```

### API Bindings Not Generated
```bash
# Run dev mode once to generate bindings
wails dev
# Press Ctrl+C after app starts
# Check frontend/wailsjs directory
```

### Service Installation Fails
```
Access denied
```
**Solution**: Run `service.bat` as Administrator
- Right-click → Run as Administrator

### Signature Verification Fails
```powershell
# Check certificate validity
certutil -verify "build\bin\FinanceHub.exe"

# View certificate details
Get-AuthenticodeSignature "build\bin\FinanceHub.exe" | Format-List
```

## 📚 Additional Resources

### Wails Documentation
- Official Docs: https://wails.io/docs/introduction
- API Reference: https://wails.io/docs/reference/runtime/intro
- Community: https://wails.io/community

### FinanceHub Documentation
- Main README: [README.md](README.md)
- Setup Guide: [SETUP.md](SETUP.md)
- Project Summary: [PROJECT_SUMMARY.md](PROJECT_SUMMARY.md)
- Quick Reference: [QUICK_REFERENCE.md](QUICK_REFERENCE.md)

### API Documentation
- Alpha Vantage: https://www.alphavantage.co/documentation/
- CoinGecko: https://www.coingecko.com/en/api/documentation

## 🆘 Getting Help

If you encounter issues:

1. Check the [Troubleshooting](#troubleshooting) section
2. Search [GitHub Issues](https://github.com/reyisjones/FinanceHub/issues)
3. Review [Wails Documentation](https://wails.io/docs/introduction)
4. Create a new issue with:
   - Error message
   - Steps to reproduce
   - System information
   - Logs from `logs/financehub_service.log`

## 🎯 Quick Command Reference

```bash
# Development
wails dev                    # Start dev server with hot reload
wails doctor                 # Check system dependencies

# Building
wails build                  # Build production executable
wails build -debug           # Build with debug console
wails build -clean           # Clean build (recommended)

# Service Management
service.bat install          # Install as Windows service
service.bat start            # Start service
service.bat stop             # Stop service
service.bat restart          # Restart service
service.bat uninstall        # Remove service

# Deployment
git tag v1.0.0              # Create release tag
git push origin v1.0.0      # Trigger GitHub Actions build
```

---

**Built with ❤️ using Wails, Go, and Material UI**
