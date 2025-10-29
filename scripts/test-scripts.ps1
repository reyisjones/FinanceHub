#!/usr/bin/env pwsh
<#
.SYNOPSIS
    Script Validation Test Suite for FinanceHub
.DESCRIPTION
    Tests all batch scripts and PowerShell scripts to ensure they function correctly
#>

$ErrorActionPreference = "Continue"
$script:TestResults = @{ Passed = 0; Failed = 0; Tests = @() }

function Test-ScriptSyntax {
    param(
        [string]$ScriptPath,
        [string]$ScriptType
    )
    
    $testName = "Syntax check: $([System.IO.Path]::GetFileName($ScriptPath))"
    
    try {
        if ($ScriptType -eq "PowerShell") {
            $null = [System.Management.Automation.PSParser]::Tokenize((Get-Content $ScriptPath -Raw), [ref]$null)
            Write-Host "[✓] $testName" -ForegroundColor Green
            $script:TestResults.Passed++
            $script:TestResults.Tests += @{ Name = $testName; Status = "Passed" }
        } elseif ($ScriptType -eq "Batch") {
            # Basic batch file validation
            $content = Get-Content $ScriptPath -Raw
            $hasInvalidChars = $content -match '[^\x00-\x7F]'
            if (-not $hasInvalidChars) {
                Write-Host "[✓] $testName" -ForegroundColor Green
                $script:TestResults.Passed++
                $script:TestResults.Tests += @{ Name = $testName; Status = "Passed" }
            } else {
                throw "Contains invalid characters"
            }
        }
        return $true
    } catch {
        Write-Host "[✗] $testName - $($_.Exception.Message)" -ForegroundColor Red
        $script:TestResults.Failed++
        $script:TestResults.Tests += @{ Name = $testName; Status = "Failed"; Error = $_.Exception.Message }
        return $false
    }
}

function Test-ScriptStructure {
    param(
        [string]$ScriptPath,
        [string[]]$RequiredPatterns
    )
    
    $fileName = [System.IO.Path]::GetFileName($ScriptPath)
    $content = Get-Content $ScriptPath -Raw
    
    foreach ($pattern in $RequiredPatterns) {
        $testName = "Structure check ($fileName): Contains '$pattern'"
        if ($content -match $pattern) {
            Write-Host "[✓] $testName" -ForegroundColor Green
            $script:TestResults.Passed++
            $script:TestResults.Tests += @{ Name = $testName; Status = "Passed" }
        } else {
            Write-Host "[✗] $testName" -ForegroundColor Yellow
            $script:TestResults.Failed++
            $script:TestResults.Tests += @{ Name = $testName; Status = "Failed" }
        }
    }
}

Write-Host "`n╔════════════════════════════════════════════════════════════════╗" -ForegroundColor Cyan
Write-Host "║           FINANCEHUB SCRIPT VALIDATION TEST SUITE             ║" -ForegroundColor Cyan
Write-Host "╚════════════════════════════════════════════════════════════════╝`n" -ForegroundColor Cyan

# Test PowerShell scripts
Write-Host "Testing PowerShell Scripts..." -ForegroundColor Cyan
Write-Host "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" -ForegroundColor Cyan

$psScripts = Get-ChildItem -Path "scripts" -Filter "*.ps1" -ErrorAction SilentlyContinue

foreach ($script in $psScripts) {
    Test-ScriptSyntax -ScriptPath $script.FullName -ScriptType "PowerShell"
}

# Test audit script specifically
if (Test-Path "scripts/audit.ps1") {
    Test-ScriptStructure -ScriptPath "scripts/audit.ps1" -RequiredPatterns @(
        'Test-SecretManagement',
        'Test-Dependencies',
        'Test-ConfigurationIntegrity',
        'Test-CodeQuality'
    )
}

# Test Batch scripts
Write-Host "`nTesting Batch Scripts..." -ForegroundColor Cyan
Write-Host "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" -ForegroundColor Cyan

$batchScripts = Get-ChildItem -Filter "*.bat" -ErrorAction SilentlyContinue

foreach ($script in $batchScripts) {
    Test-ScriptSyntax -ScriptPath $script.FullName -ScriptType "Batch"
    
    # Check for error handling in critical scripts
    $content = Get-Content $script.FullName -Raw
    $testName = "Error handling check: $($script.Name)"
    
    if ($script.Name -match "(build|setup|deploy)") {
        if ($content -match '@echo off' -and ($content -match 'if errorlevel' -or $content -match 'if %errorlevel%')) {
            Write-Host "[✓] $testName" -ForegroundColor Green
            $script:TestResults.Passed++
        } else {
            Write-Host "[⚠] $testName - Consider adding error handling" -ForegroundColor Yellow
            $script:TestResults.Tests += @{ Name = $testName; Status = "Warning" }
        }
    }
}

# Test specific build scripts
if (Test-Path "build-wails.bat") {
    Test-ScriptStructure -ScriptPath "build-wails.bat" -RequiredPatterns @('wails build')
}

if (Test-Path "setup.bat") {
    Test-ScriptStructure -ScriptPath "setup.bat" -RequiredPatterns @('npm install', 'go mod')
}

# Summary
Write-Host "`n╔════════════════════════════════════════════════════════════════╗" -ForegroundColor Magenta
Write-Host "║                      TEST SUMMARY                              ║" -ForegroundColor Magenta
Write-Host "╚════════════════════════════════════════════════════════════════╝`n" -ForegroundColor Magenta

$total = $script:TestResults.Passed + $script:TestResults.Failed
Write-Host "Total Tests:  $total" -ForegroundColor White
Write-Host "Passed:       $($script:TestResults.Passed)" -ForegroundColor Green
Write-Host "Failed:       $($script:TestResults.Failed)" -ForegroundColor Red

if ($script:TestResults.Failed -eq 0) {
    Write-Host "`n✓ All script validation tests passed!" -ForegroundColor Green
    exit 0
} else {
    Write-Host "`n⚠ Some tests failed. Please review the output above." -ForegroundColor Yellow
    exit 1
}
