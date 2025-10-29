#!/usr/bin/env pwsh
<#
.SYNOPSIS
    Security and Testing Compliance Audit Script for FinanceHub
.DESCRIPTION
    Comprehensive audit script that validates:
    - Configuration integrity
    - Secret management and sensitive data exposure
    - Dependency vulnerabilities
    - Code quality and test coverage
    - Security best practices
.NOTES
    Version: 1.0.0
    Author: FinanceHub Team
#>

param(
    [switch]$Verbose,
    [switch]$FixIssues,
    [string]$OutputFormat = "console" # console, json, html
)

$ErrorActionPreference = "Stop"
$script:AuditResults = @{
    Timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
    Passed = @()
    Failed = @()
    Warnings = @()
    Score = 0
}

function Write-AuditHeader {
    param([string]$Message)
    Write-Host "`n========================================" -ForegroundColor Cyan
    Write-Host "  $Message" -ForegroundColor Cyan
    Write-Host "========================================`n" -ForegroundColor Cyan
}

function Write-AuditResult {
    param(
        [string]$Test,
        [bool]$Passed,
        [string]$Message,
        [string]$Severity = "Info"
    )
    
    if ($Passed) {
        Write-Host "[✓] $Test" -ForegroundColor Green
        $script:AuditResults.Passed += @{ Test = $Test; Message = $Message }
        $script:AuditResults.Score += 10
    } else {
        $color = if ($Severity -eq "Critical") { "Red" } elseif ($Severity -eq "Warning") { "Yellow" } else { "Red" }
        Write-Host "[✗] $Test" -ForegroundColor $color
        Write-Host "    $Message" -ForegroundColor $color
        
        if ($Severity -eq "Warning") {
            $script:AuditResults.Warnings += @{ Test = $Test; Message = $Message }
            $script:AuditResults.Score += 5
        } else {
            $script:AuditResults.Failed += @{ Test = $Test; Message = $Message; Severity = $Severity }
        }
    }
}

# =============================================================================
# 1. SECRET MANAGEMENT AUDIT
# =============================================================================
function Test-SecretManagement {
    Write-AuditHeader "SECRET MANAGEMENT AUDIT"
    
    # Check for .env files in git
    $envInGit = git ls-files | Select-String -Pattern "\.env$" -Quiet
    Write-AuditResult `
        -Test "No .env files in git repository" `
        -Passed (-not $envInGit) `
        -Message "Found .env files in git. These should be in .gitignore" `
        -Severity "Critical"
    
    # Check .gitignore exists and contains common secrets
    $gitignorePath = ".gitignore"
    if (Test-Path $gitignorePath) {
        $gitignoreContent = Get-Content $gitignorePath -Raw
        $hasEnvPattern = $gitignoreContent -match "\.env"
        $hasKeysPattern = $gitignoreContent -match "(\*\.key|\*\.pem|\*\.pfx)"
        
        Write-AuditResult `
            -Test ".gitignore excludes .env files" `
            -Passed $hasEnvPattern `
            -Message ".gitignore should exclude .env files" `
            -Severity "Critical"
            
        Write-AuditResult `
            -Test ".gitignore excludes key/certificate files" `
            -Passed $hasKeysPattern `
            -Message ".gitignore should exclude *.key, *.pem, *.pfx files" `
            -Severity "Warning"
    }
    
    # Scan for hardcoded secrets in source files
    $secretPatterns = @(
        @{ Pattern = 'password\s*=\s*[''"](?!.*\{\{)[^''"]+[''"]'; Name = "Hardcoded passwords" }
        @{ Pattern = 'api[_-]?key\s*=\s*[''"](?!.*\{\{)[^''"]+[''"]'; Name = "Hardcoded API keys" }
        @{ Pattern = 'secret\s*=\s*[''"](?!.*\{\{)[^''"]+[''"]'; Name = "Hardcoded secrets" }
        @{ Pattern = '(?i)(ghp|github_pat)_[a-zA-Z0-9]{36,}'; Name = "GitHub tokens" }
    )
    
    $foundSecrets = $false
    foreach ($pattern in $secretPatterns) {
        $matches = Get-ChildItem -Recurse -Include *.go,*.ts,*.tsx,*.js,*.json,*.yml,*.yaml `
            | Select-String -Pattern $pattern.Pattern -CaseSensitive:$false
        
        if ($matches) {
            $foundSecrets = $true
            Write-Host "  Found potential $($pattern.Name):" -ForegroundColor Yellow
            $matches | ForEach-Object { Write-Host "    $($_.Path):$($_.LineNumber)" -ForegroundColor Yellow }
        }
    }
    
    Write-AuditResult `
        -Test "No hardcoded secrets in source code" `
        -Passed (-not $foundSecrets) `
        -Message "Found potential hardcoded secrets. Review and move to environment variables" `
        -Severity "Critical"
}

# =============================================================================
# 2. DEPENDENCY VULNERABILITY AUDIT
# =============================================================================
function Test-Dependencies {
    Write-AuditHeader "DEPENDENCY VULNERABILITY AUDIT"
    
    # Go dependencies
    if (Test-Path "go.mod") {
        Write-Host "Checking Go dependencies..." -ForegroundColor Cyan
        
        # Check for go.sum
        Write-AuditResult `
            -Test "go.sum file exists" `
            -Passed (Test-Path "go.sum") `
            -Message "go.sum is required for reproducible builds" `
            -Severity "Critical"
        
        # Run go mod verify
        try {
            $verifyOutput = go mod verify 2>&1
            $verified = $LASTEXITCODE -eq 0
            Write-AuditResult `
                -Test "Go modules verified" `
                -Passed $verified `
                -Message "Go module verification failed: $verifyOutput" `
                -Severity "Critical"
        } catch {
            Write-AuditResult `
                -Test "Go modules verified" `
                -Passed $false `
                -Message "Failed to run go mod verify" `
                -Severity "Critical"
        }
        
        # Check for known vulnerable packages (example)
        $vulnPackages = @("github.com/go-sql-driver/mysql@v1.4.0")
        $goModContent = Get-Content "go.mod" -Raw
        $hasVuln = $false
        foreach ($pkg in $vulnPackages) {
            if ($goModContent -match [regex]::Escape($pkg)) {
                $hasVuln = $true
                Write-Host "  Found vulnerable package: $pkg" -ForegroundColor Yellow
            }
        }
        Write-AuditResult `
            -Test "No known vulnerable Go packages" `
            -Passed (-not $hasVuln) `
            -Message "Update vulnerable packages" `
            -Severity "Critical"
    }
    
    # Node.js dependencies
    if (Test-Path "frontend/package.json") {
        Write-Host "Checking Node.js dependencies..." -ForegroundColor Cyan
        
        # Check package-lock.json exists
        Write-AuditResult `
            -Test "package-lock.json exists" `
            -Passed (Test-Path "frontend/package-lock.json") `
            -Message "package-lock.json ensures reproducible builds" `
            -Severity "Warning"
        
        # Run npm audit (if npm is available)
        try {
            Push-Location "frontend"
            if (Get-Command npm -ErrorAction SilentlyContinue) {
                $auditOutput = npm audit --json 2>&1 | ConvertFrom-Json
                $criticalVulns = $auditOutput.metadata.vulnerabilities.critical
                $highVulns = $auditOutput.metadata.vulnerabilities.high
                
                Write-AuditResult `
                    -Test "No critical npm vulnerabilities" `
                    -Passed ($criticalVulns -eq 0) `
                    -Message "Found $criticalVulns critical and $highVulns high severity vulnerabilities" `
                    -Severity "Critical"
            }
            Pop-Location
        } catch {
            Write-AuditResult `
                -Test "npm audit completed" `
                -Passed $false `
                -Message "npm audit failed or npm not installed" `
                -Severity "Warning"
            Pop-Location
        }
    }
}

# =============================================================================
# 3. CONFIGURATION INTEGRITY AUDIT
# =============================================================================
function Test-ConfigurationIntegrity {
    Write-AuditHeader "CONFIGURATION INTEGRITY AUDIT"
    
    # Check required configuration files
    $requiredConfigs = @(
        @{ Path = "wails.json"; Name = "Wails configuration" }
        @{ Path = "go.mod"; Name = "Go module file" }
        @{ Path = "frontend/package.json"; Name = "Frontend package.json" }
        @{ Path = ".gitignore"; Name = "Git ignore file" }
    )
    
    foreach ($config in $requiredConfigs) {
        Write-AuditResult `
            -Test "$($config.Name) exists" `
            -Passed (Test-Path $config.Path) `
            -Message "$($config.Path) is missing" `
            -Severity "Critical"
    }
    
    # Validate wails.json structure
    if (Test-Path "wails.json") {
        try {
            $wailsConfig = Get-Content "wails.json" | ConvertFrom-Json
            $hasName = $null -ne $wailsConfig.name
            $hasVersion = $null -ne $wailsConfig.info.productVersion
            $hasFrontendBuild = $null -ne $wailsConfig.frontend.'build'
            
            Write-AuditResult `
                -Test "wails.json is valid JSON" `
                -Passed $true `
                -Message "wails.json is valid"
            
            Write-AuditResult `
                -Test "wails.json has required fields" `
                -Passed ($hasName -and $hasVersion -and $hasFrontendBuild) `
                -Message "wails.json missing required fields" `
                -Severity "Critical"
        } catch {
            Write-AuditResult `
                -Test "wails.json is valid JSON" `
                -Passed $false `
                -Message "wails.json is not valid JSON" `
                -Severity "Critical"
        }
    }
    
    # Check VERSION file
    if (Test-Path "VERSION") {
        $version = Get-Content "VERSION" -Raw
        $validVersion = $version -match '^\d+\.\d+\.\d+$'
        Write-AuditResult `
            -Test "VERSION file has valid semver format" `
            -Passed $validVersion `
            -Message "VERSION should follow semantic versioning (e.g., 1.0.0)" `
            -Severity "Warning"
    }
}

# =============================================================================
# 4. CODE QUALITY & TEST COVERAGE AUDIT
# =============================================================================
function Test-CodeQuality {
    Write-AuditHeader "CODE QUALITY & TEST COVERAGE AUDIT"
    
    # Check for test files
    $goTestFiles = Get-ChildItem -Recurse -Filter "*_test.go" -ErrorAction SilentlyContinue
    $tsTestFiles = Get-ChildItem -Recurse -Filter "*.test.ts*" -ErrorAction SilentlyContinue
    
    Write-AuditResult `
        -Test "Go test files exist" `
        -Passed ($goTestFiles.Count -gt 0) `
        -Message "No Go test files found. Add *_test.go files" `
        -Severity "Warning"
    
    Write-AuditResult `
        -Test "TypeScript test files exist" `
        -Passed ($tsTestFiles.Count -gt 0) `
        -Message "No TypeScript test files found. Add *.test.ts files" `
        -Severity "Warning"
    
    # Run Go tests if available
    if ($goTestFiles.Count -gt 0) {
        try {
            Write-Host "Running Go tests..." -ForegroundColor Cyan
            $testOutput = go test ./... -v 2>&1
            $testsPassed = $LASTEXITCODE -eq 0
            
            Write-AuditResult `
                -Test "Go tests pass" `
                -Passed $testsPassed `
                -Message "Go tests failed. Fix failing tests" `
                -Severity "Critical"
            
            if ($Verbose) {
                Write-Host $testOutput
            }
        } catch {
            Write-AuditResult `
                -Test "Go tests executed" `
                -Passed $false `
                -Message "Failed to run Go tests" `
                -Severity "Warning"
        }
    }
    
    # Check for linter configuration
    Write-AuditResult `
        -Test "Frontend ESLint configured" `
        -Passed (Test-Path "frontend/.eslintrc.cjs" -or Test-Path "frontend/.eslintrc.js") `
        -Message "ESLint configuration not found" `
        -Severity "Warning"
}

# =============================================================================
# 5. SECURITY BEST PRACTICES AUDIT
# =============================================================================
function Test-SecurityBestPractices {
    Write-AuditHeader "SECURITY BEST PRACTICES AUDIT"
    
    # Check for HTTPS in production
    $backendMain = Get-Content "backend/main.go" -Raw -ErrorAction SilentlyContinue
    if ($backendMain) {
        $hasHTTPS = $backendMain -match 'ListenAndServeTLS|https://'
        Write-AuditResult `
            -Test "Backend configured for HTTPS" `
            -Passed $hasHTTPS `
            -Message "Consider using HTTPS in production" `
            -Severity "Warning"
    }
    
    # Check CORS configuration
    if ($backendMain) {
        $hasCORS = $backendMain -match 'cors'
        Write-AuditResult `
            -Test "CORS middleware configured" `
            -Passed $hasCORS `
            -Message "CORS configuration not found" `
            -Severity "Warning"
    }
    
    # Check for security headers
    $hasSecurityHeaders = $false
    if ($backendMain) {
        $securityHeaders = @(
            'X-Frame-Options',
            'X-Content-Type-Options',
            'X-XSS-Protection'
        )
        foreach ($header in $securityHeaders) {
            if ($backendMain -match $header) {
                $hasSecurityHeaders = $true
                break
            }
        }
    }
    Write-AuditResult `
        -Test "Security headers configured" `
        -Passed $hasSecurityHeaders `
        -Message "Add security headers (X-Frame-Options, X-Content-Type-Options, etc.)" `
        -Severity "Warning"
    
    # Check for SQL injection protection (parameterized queries)
    $goFiles = Get-ChildItem -Recurse -Filter "*.go" -ErrorAction SilentlyContinue
    $hasSQLInjectionRisk = $false
    foreach ($file in $goFiles) {
        $content = Get-Content $file.FullName -Raw
        if ($content -match 'Query\s*\(\s*fmt\.Sprintf|Exec\s*\(\s*fmt\.Sprintf') {
            $hasSQLInjectionRisk = $true
            Write-Host "  Potential SQL injection risk in: $($file.Name)" -ForegroundColor Yellow
        }
    }
    Write-AuditResult `
        -Test "No SQL injection vulnerabilities" `
        -Passed (-not $hasSQLInjectionRisk) `
        -Message "Use parameterized queries instead of string concatenation" `
        -Severity "Critical"
}

# =============================================================================
# 6. GITHUB ACTIONS SECURITY AUDIT
# =============================================================================
function Test-GitHubActionsSecurity {
    Write-AuditHeader "CI/CD SECURITY AUDIT"
    
    $workflowFiles = Get-ChildItem -Path ".github/workflows" -Filter "*.yml" -ErrorAction SilentlyContinue
    
    if ($workflowFiles) {
        foreach ($workflow in $workflowFiles) {
            $content = Get-Content $workflow.FullName -Raw
            
            # Check for hardcoded credentials
            $hasHardcodedCreds = $content -match 'password:\s*[''"](?!.*\$\{\{)[^''"]+[''"]'
            Write-AuditResult `
                -Test "[$($workflow.Name)] No hardcoded credentials" `
                -Passed (-not $hasHardcodedCreds) `
                -Message "Use GitHub Secrets for credentials" `
                -Severity "Critical"
            
            # Check for pinned action versions
            $hasUnpinnedActions = $content -match 'uses:.*@(main|master|latest)'
            Write-AuditResult `
                -Test "[$($workflow.Name)] Actions pinned to specific versions" `
                -Passed (-not $hasUnpinnedActions) `
                -Message "Pin actions to specific SHA or version tags" `
                -Severity "Warning"
        }
    }
}

# =============================================================================
# MAIN EXECUTION
# =============================================================================
function Invoke-Audit {
    Write-Host "`n" -NoNewline
    Write-Host "╔════════════════════════════════════════════════════════════════╗" -ForegroundColor Magenta
    Write-Host "║                                                                ║" -ForegroundColor Magenta
    Write-Host "║         FINANCEHUB SECURITY & COMPLIANCE AUDIT v1.0.0         ║" -ForegroundColor Magenta
    Write-Host "║                                                                ║" -ForegroundColor Magenta
    Write-Host "╚════════════════════════════════════════════════════════════════╝" -ForegroundColor Magenta
    Write-Host "`n"
    
    Test-SecretManagement
    Test-Dependencies
    Test-ConfigurationIntegrity
    Test-CodeQuality
    Test-SecurityBestPractices
    Test-GitHubActionsSecurity
    
    # Generate summary
    Write-AuditHeader "AUDIT SUMMARY"
    
    $totalTests = $script:AuditResults.Passed.Count + $script:AuditResults.Failed.Count + $script:AuditResults.Warnings.Count
    $passRate = if ($totalTests -gt 0) { [math]::Round(($script:AuditResults.Passed.Count / $totalTests) * 100, 2) } else { 0 }
    
    Write-Host "Total Tests Run:    $totalTests" -ForegroundColor White
    Write-Host "Passed:            $($script:AuditResults.Passed.Count)" -ForegroundColor Green
    Write-Host "Failed:            $($script:AuditResults.Failed.Count)" -ForegroundColor Red
    Write-Host "Warnings:          $($script:AuditResults.Warnings.Count)" -ForegroundColor Yellow
    Write-Host "Pass Rate:         $passRate%" -ForegroundColor $(if ($passRate -ge 80) { "Green" } elseif ($passRate -ge 60) { "Yellow" } else { "Red" })
    Write-Host "Security Score:    $($script:AuditResults.Score)/$(($totalTests * 10))" -ForegroundColor Cyan
    
    if ($script:AuditResults.Failed.Count -gt 0) {
        Write-Host "`nCritical Issues:" -ForegroundColor Red
        foreach ($failure in $script:AuditResults.Failed) {
            if ($failure.Severity -eq "Critical") {
                Write-Host "  • $($failure.Test): $($failure.Message)" -ForegroundColor Red
            }
        }
    }
    
    # Export results if requested
    if ($OutputFormat -eq "json") {
        $jsonOutput = $script:AuditResults | ConvertTo-Json -Depth 10
        $jsonOutput | Out-File "audit-results.json"
        Write-Host "`nResults exported to audit-results.json" -ForegroundColor Cyan
    }
    
    Write-Host "`n"
    
    # Exit with error if critical failures
    if ($script:AuditResults.Failed.Count -gt 0) {
        $criticalCount = ($script:AuditResults.Failed | Where-Object { $_.Severity -eq "Critical" }).Count
        if ($criticalCount -gt 0) {
            Write-Host "⚠ Audit completed with $criticalCount critical issue(s). Please review and fix." -ForegroundColor Red
            exit 1
        }
    } else {
        Write-Host "✓ Audit completed successfully! All checks passed." -ForegroundColor Green
    }
}

# Run the audit
Invoke-Audit
