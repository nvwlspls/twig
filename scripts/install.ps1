# Twig installer script for Windows
# Usage: Invoke-Expression (Invoke-WebRequest -Uri "https://raw.githubusercontent.com/yourusername/twig/main/scripts/install.ps1" -UseBasicParsing).Content

param(
    [string]$Version = "1.0.0",
    [string]$InstallPath = "$env:USERPROFILE\bin"
)

$ErrorActionPreference = 'Stop'

$Repo = "yourusername/twig"
$BinaryName = "twig"

# Detect architecture
$Arch = if ([Environment]::Is64BitOperatingSystem) { "amd64" } else { "386" }
$OS = "windows"

# Download URL
$DownloadUrl = "https://github.com/$Repo/releases/download/v$Version/${BinaryName}-${Version}-${OS}-${Arch}.tar.gz"

Write-Host "Installing twig v$Version for $OS/$Arch..." -ForegroundColor Green

# Create temporary directory
$TempDir = New-TemporaryFile | ForEach-Object { Remove-Item $_; New-Item -ItemType Directory -Path $_ }
$TarFile = Join-Path $TempDir.FullName "twig.tar.gz"

try {
    # Download
    Write-Host "Downloading from $DownloadUrl..." -ForegroundColor Yellow
    Invoke-WebRequest -Uri $DownloadUrl -OutFile $TarFile

    # Extract (requires tar command, available in Windows 10 1803+)
    Write-Host "Extracting..." -ForegroundColor Yellow
    tar -xzf $TarFile -C $TempDir.FullName

    # Create install directory
    if (!(Test-Path $InstallPath)) {
        New-Item -ItemType Directory -Path $InstallPath -Force | Out-Null
    }

    # Move binary
    $BinaryFile = Join-Path $TempDir.FullName "twig.exe"
    if (Test-Path $BinaryFile) {
        Copy-Item $BinaryFile $InstallPath -Force
        Write-Host "Binary installed to $InstallPath" -ForegroundColor Green
    } else {
        throw "Executable not found in extracted archive"
    }

    # Add to PATH if not already there
    $CurrentPath = [Environment]::GetEnvironmentVariable("PATH", "User")
    if ($CurrentPath -notlike "*$InstallPath*") {
        $NewPath = "$CurrentPath;$InstallPath"
        [Environment]::SetEnvironmentVariable("PATH", $NewPath, "User")
        Write-Host "Added $InstallPath to PATH" -ForegroundColor Green
    }

    Write-Host "✅ twig v$Version installed successfully!" -ForegroundColor Green
    Write-Host ""
    Write-Host "To use twig, either:" -ForegroundColor Yellow
    Write-Host "1. Restart your terminal, or" -ForegroundColor Yellow
    Write-Host "2. Refresh your PATH: `$env:PATH = [Environment]::GetEnvironmentVariable('PATH', 'User')" -ForegroundColor Yellow
    Write-Host ""
    Write-Host "Then run: twig --version" -ForegroundColor Yellow

} catch {
    Write-Host "❌ Installation failed: $($_.Exception.Message)" -ForegroundColor Red
    exit 1
} finally {
    # Clean up
    if (Test-Path $TempDir.FullName) {
        Remove-Item $TempDir.FullName -Recurse -Force
    }
} 