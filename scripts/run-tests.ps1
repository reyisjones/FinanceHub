#!/usr/bin/env pwsh
<#
.SYNOPSIS
    Master Test Runner for FinanceHub
.DESCRIPTION
    Runs all tests across the project: frontend, backend, scripts, and audit
.PARAMETER Suite
    Test suite to run: all, frontend, backend, scripts, audit, ci
.PARAMETER Coverage
    Generate coverage reports
#>

param(
    [ValidateSet("all", "frontend", "backend", "scripts", "audit", "ci")]
    [string]$Suite = "all",
    [switch]$Coverage,
    [switch]$Verbose
)

$ErrorActionPreference = "Continue"
$script:TestSummary = @{
    Frontend = @{ Status = "Not Run"; Passed = 0; Failed = 0 }
    Backend = @{ Status = "Not Run"; Passed = 0; Failed = 0 }
    Scripts = @{ Status = "Not Run"; Passed = 0; Failed = 0 }
    Audit = @{ Status = "Not Run"; Passed = 0; Failed = 0 }
}

function Write-TestHeader {
    param([string]$Message)
    Write-Host "`n" -NoNewline
    Write-Host "╔════════════════════════════════════════════════════════════════╗" -ForegroundColor Cyan
    Write-Host "║  $($Message.PadRight(62))║" -ForegroundColor Cyan
    Write-Host "╚════════════════════════════════════════════════════════════════╝" -ForegroundColor Cyan
    Write-Host ""
}

function Test-Frontend {
    Write-TestHeader "FRONTEND TESTS (TypeScript/React)"
    
    if (-not (Test-Path "frontend/package.json")) {
        Write-Host "⚠ Frontend package.json not found. Skipping frontend tests." -ForegroundColor Yellow
        return
    }
    
    Push-Location "frontend"
    
    try {
        # Check if test dependencies are installed
        if (-not (Test-Path "node_modules")) {
            Write-Host "Installing frontend dependencies..." -ForegroundColor Cyan
            npm install
        }
        
        # Check if vitest is configured
        if (Test-Path "vitest.config.ts") {
            Write-Host "Running frontend tests with Vitest..." -ForegroundColor Cyan
            
            if ($Coverage) {
                npm run test -- --coverage
            } else {
                npm run test
            }
            
            if ($LASTEXITCODE -eq 0) {
                $script:TestSummary.Frontend.Status = "Passed"
                $script:TestSummary.Frontend.Passed = 1
                Write-Host "✓ Frontend tests passed" -ForegroundColor Green
            } else {
                $script:TestSummary.Frontend.Status = "Failed"
                $script:TestSummary.Frontend.Failed = 1
                Write-Host "✗ Frontend tests failed" -ForegroundColor Red
            }
        } else {
            Write-Host "⚠ Vitest not configured. Add vitest.config.ts to run tests." -ForegroundColor Yellow
            Write-Host "  Test files found but cannot execute without test runner." -ForegroundColor Yellow
            $script:TestSummary.Frontend.Status = "Skipped"
        }
    } catch {
        Write-Host "✗ Error running frontend tests: $($_.Exception.Message)" -ForegroundColor Red
        $script:TestSummary.Frontend.Status = "Error"
        $script:TestSummary.Frontend.Failed = 1
    } finally {
        Pop-Location
    }
}

function Test-Backend {
    Write-TestHeader "BACKEND TESTS (Go)"
    
    if (-not (Test-Path "go.mod")) {
        Write-Host "⚠ go.mod not found. Skipping backend tests." -ForegroundColor Yellow
        return
    }
    
    try {
        # Download dependencies
        Write-Host "Downloading Go dependencies..." -ForegroundColor Cyan
        go mod download
        
        # Run tests
        Write-Host "Running Go tests..." -ForegroundColor Cyan
        
        if ($Coverage) {
            go test ./... -v -coverprofile=coverage.out
            if ($LASTEXITCODE -eq 0) {
                Write-Host "`nGenerating coverage report..." -ForegroundColor Cyan
                go tool cover -html=coverage.out -o coverage.html
                Write-Host "Coverage report saved to coverage.html" -ForegroundColor Cyan
            }
        } else {
            go test ./... -v
        }
        
        if ($LASTEXITCODE -eq 0) {
            $script:TestSummary.Backend.Status = "Passed"
            $script:TestSummary.Backend.Passed = 1
            Write-Host "✓ Backend tests passed" -ForegroundColor Green
        } else {
            $script:TestSummary.Backend.Status = "Failed"
            $script:TestSummary.Backend.Failed = 1
            Write-Host "✗ Backend tests failed" -ForegroundColor Red
        }
    } catch {
        Write-Host "✗ Error running backend tests: $($_.Exception.Message)" -ForegroundColor Red
        $script:TestSummary.Backend.Status = "Error"
        $script:TestSummary.Backend.Failed = 1
    }
}

function Test-Scripts {
    Write-TestHeader "SCRIPT VALIDATION TESTS"
    
    if (Test-Path "scripts/test-scripts.ps1") {
        try {
            & "scripts/test-scripts.ps1"
            
            if ($LASTEXITCODE -eq 0) {
                $script:TestSummary.Scripts.Status = "Passed"
                $script:TestSummary.Scripts.Passed = 1
            } else {
                $script:TestSummary.Scripts.Status = "Failed"
                $script:TestSummary.Scripts.Failed = 1
            }
        } catch {
            Write-Host "✗ Error running script tests: $($_.Exception.Message)" -ForegroundColor Red
            $script:TestSummary.Scripts.Status = "Error"
            $script:TestSummary.Scripts.Failed = 1
        }
    } else {
        Write-Host "⚠ Script test suite not found at scripts/test-scripts.ps1" -ForegroundColor Yellow
        $script:TestSummary.Scripts.Status = "Skipped"
    }
}

function Test-SecurityAudit {
    Write-TestHeader "SECURITY & COMPLIANCE AUDIT"
    
    if (Test-Path "scripts/audit.ps1") {
        try {
            $auditArgs = @()
            if ($Verbose) { $auditArgs += "-Verbose" }
            
            & "scripts/audit.ps1" @auditArgs
            
            if ($LASTEXITCODE -eq 0) {
                $script:TestSummary.Audit.Status = "Passed"
                $script:TestSummary.Audit.Passed = 1
            } else {
                $script:TestSummary.Audit.Status = "Failed"
                $script:TestSummary.Audit.Failed = 1
            }
        } catch {
            Write-Host "✗ Error running security audit: $($_.Exception.Message)" -ForegroundColor Red
            $script:TestSummary.Audit.Status = "Error"
            $script:TestSummary.Audit.Failed = 1
        }
    } else {
        Write-Host "⚠ Audit script not found at scripts/audit.ps1" -ForegroundColor Yellow
        $script:TestSummary.Audit.Status = "Skipped"
    }
}

function Show-Summary {
    Write-Host "`n" -NoNewline
    Write-Host "╔════════════════════════════════════════════════════════════════╗" -ForegroundColor Magenta
    Write-Host "║                      OVERALL TEST SUMMARY                      ║" -ForegroundColor Magenta
    Write-Host "╚════════════════════════════════════════════════════════════════╝" -ForegroundColor Magenta
    Write-Host ""
    
    $totalPassed = 0
    $totalFailed = 0
    
    foreach ($suite in $script:TestSummary.Keys) {
        $result = $script:TestSummary[$suite]
        $status = $result.Status
        $color = switch ($status) {
            "Passed" { "Green" }
            "Failed" { "Red" }
            "Error" { "Red" }
            "Skipped" { "Yellow" }
            default { "Gray" }
        }
        
        $statusSymbol = switch ($status) {
            "Passed" { "✓" }
            "Failed" { "✗" }
            "Error" { "⚠" }
            "Skipped" { "○" }
            default { "-" }
        }
        
        Write-Host "  $statusSymbol " -NoNewline -ForegroundColor $color
        Write-Host "$($suite.PadRight(15))" -NoNewline
        Write-Host " $status" -ForegroundColor $color
        
        $totalPassed += $result.Passed
        $totalFailed += $result.Failed
    }
    
    Write-Host ""
    Write-Host "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" -ForegroundColor Gray
    Write-Host "  Total Passed: $totalPassed" -ForegroundColor Green
    Write-Host "  Total Failed: $totalFailed" -ForegroundColor Red
    Write-Host ""
    
    if ($totalFailed -eq 0) {
        Write-Host "  🎉 All test suites passed!" -ForegroundColor Green
        Write-Host ""
        exit 0
    } else {
        Write-Host "  ⚠  Some test suites failed. Please review the output above." -ForegroundColor Yellow
        Write-Host ""
        exit 1
    }
}

# Main execution
Write-Host "`n" -NoNewline
Write-Host "╔════════════════════════════════════════════════════════════════╗" -ForegroundColor Magenta
Write-Host "║                                                                ║" -ForegroundColor Magenta
Write-Host "║              FINANCEHUB COMPREHENSIVE TEST SUITE               ║" -ForegroundColor Magenta
Write-Host "║                        Version 1.0.0                           ║" -ForegroundColor Magenta
Write-Host "║                                                                ║" -ForegroundColor Magenta
Write-Host "╚════════════════════════════════════════════════════════════════╝" -ForegroundColor Magenta
Write-Host ""
Write-Host "Running test suite: $Suite" -ForegroundColor Cyan
if ($Coverage) {
    Write-Host "Coverage reporting: Enabled" -ForegroundColor Cyan
}
Write-Host ""

switch ($Suite) {
    "all" {
        Test-Backend
        Test-Frontend
        Test-Scripts
        Test-SecurityAudit
    }
    "frontend" {
        Test-Frontend
    }
    "backend" {
        Test-Backend
    }
    "scripts" {
        Test-Scripts
    }
    "audit" {
        Test-SecurityAudit
    }
    "ci" {
        # CI mode runs critical tests only
        Test-Backend
        Test-SecurityAudit
    }
}

Show-Summary
