$ErrorActionPreference = 'Stop'

$packageName = 'twig'
$url = 'https://github.com/yourusername/twig/releases/download/v1.0.0/twig-1.0.0-windows-amd64.tar.gz'
$checksum = 'PLACEHOLDER_SHA256'
$checksumType = 'sha256'
$toolsDir = "$(Split-Path -parent $MyInvocation.MyCommand.Definition)"

# Download and extract
$tempDir = Join-Path $env:TEMP $packageName
if (!(Test-Path $tempDir)) { New-Item -ItemType Directory -Path $tempDir | Out-Null }

$tarFile = Join-Path $tempDir "twig.tar.gz"
Invoke-WebRequest -Uri $url -OutFile $tarFile

# Verify checksum
$hash = Get-FileHash -Path $tarFile -Algorithm $checksumType
if ($hash.Hash -ne $checksum) {
    throw "Checksum verification failed. Expected: $checksum, Got: $($hash.Hash)"
}

# Extract
tar -xzf $tarFile -C $tempDir

# Install
$exeFile = Join-Path $tempDir "twig.exe"
if (Test-Path $exeFile) {
    Copy-Item $exeFile $toolsDir -Force
} else {
    throw "Executable not found in extracted archive"
}

# Clean up
Remove-Item $tempDir -Recurse -Force -ErrorAction SilentlyContinue 