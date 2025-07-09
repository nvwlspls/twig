#!/bin/bash

# Twig installer script
# Usage: curl -fsSL https://raw.githubusercontent.com/yourusername/twig/main/scripts/install.sh | bash

set -e

VERSION="1.0.0"
REPO="yourusername/twig"
BINARY_NAME="twig"

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64)
        ARCH="amd64"
        ;;
    aarch64|arm64)
        ARCH="arm64"
        ;;
    i386|i686)
        ARCH="386"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

# Set file extension
if [ "$OS" = "windows" ]; then
    EXT=".exe"
else
    EXT=""
fi

# Download URL
DOWNLOAD_URL="https://github.com/$REPO/releases/download/v$VERSION/${BINARY_NAME}-${VERSION}-${OS}-${ARCH}.tar.gz"

echo "Installing twig v$VERSION for $OS/$ARCH..."

# Create temporary directory
TEMP_DIR=$(mktemp -d)
cd "$TEMP_DIR"

# Download and extract
echo "Downloading from $DOWNLOAD_URL..."
curl -fsSL "$DOWNLOAD_URL" -o "${BINARY_NAME}.tar.gz"

# Extract
tar -xzf "${BINARY_NAME}.tar.gz"

# Make executable
chmod +x "${BINARY_NAME}${EXT}"

# Install to user's bin directory
USER_BIN="$HOME/bin"
if [ ! -d "$USER_BIN" ]; then
    mkdir -p "$USER_BIN"
fi

# Move binary
mv "${BINARY_NAME}${EXT}" "$USER_BIN/"

# Add to PATH if not already there
if [[ ":$PATH:" != *":$USER_BIN:"* ]]; then
    echo "Adding $USER_BIN to PATH..."
    if [[ "$SHELL" == *"zsh"* ]]; then
        echo 'export PATH="$PATH:$HOME/bin"' >> ~/.zshrc
        echo "Added to ~/.zshrc"
    elif [[ "$SHELL" == *"bash"* ]]; then
        echo 'export PATH="$PATH:$HOME/bin"' >> ~/.bashrc
        echo "Added to ~/.bashrc"
    else
        echo "Please add $USER_BIN to your PATH manually"
    fi
fi

# Clean up
cd /
rm -rf "$TEMP_DIR"

echo "✅ twig v$VERSION installed successfully!"
echo ""
echo "To use twig, either:"
echo "1. Restart your terminal, or"
echo "2. Run: export PATH=\"\$PATH:\$HOME/bin\""
echo ""
echo "Then run: twig --version" 