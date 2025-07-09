# Installation Guide

## Quick Start

### Download Pre-built Binary

1. **Go to [GitHub Releases](https://github.com/yourusername/twig/releases)**
2. **Download the binary for your platform:**
   - **Windows**: `twig-{version}-windows-amd64.tar.gz`
   - **Linux**: `twig-{version}-linux-amd64.tar.gz`
   - **macOS**: `twig-{version}-darwin-amd64.tar.gz`

### Extract and Install

#### Windows
```powershell
# Download and extract
tar -xzf twig-1.0.0-windows-amd64.tar.gz

# Move to a directory in your PATH (optional)
move twig.exe C:\Windows\System32\
# OR add current directory to PATH
```

#### Linux/macOS
```bash
# Download and extract
tar -xzf twig-1.0.0-linux-amd64.tar.gz

# Make executable
chmod +x twig

# Move to a directory in your PATH
sudo mv twig /usr/local/bin/
# OR add to PATH in your shell profile
echo 'export PATH="$PATH:$HOME/bin"' >> ~/.bashrc
mv twig ~/bin/
```

### Verify Installation

```bash
twig --version
# Should output: twig version 1.0.0
```

## Usage Examples

### Basic Usage
```bash
# Create a simple site
mkdir my-site
cd my-site

# Create content directory
mkdir content

# Add some markdown files
echo "# Hello World" > content/index.md
echo "# About" > content/about.md

# Generate the site
twig

# Serve locally
twig --serve
```

### Custom Configuration
```bash
# Use custom directories
twig --source blog --output docs

# Use custom template
twig --template custom.html

# Serve on custom port
twig --serve --port 3000
```

## Package Managers

### Using Go Install (if Go is installed)
```bash
go install github.com/yourusername/twig@latest
```

### Using Homebrew (macOS)
```bash
# Add custom tap (if available)
brew tap yourusername/twig
brew install twig

# Or install directly from release
brew install --formula https://raw.githubusercontent.com/yourusername/twig/main/Formula/twig.rb
```

### Using Chocolatey (Windows)
```powershell
# Install from custom source (if available)
choco install twig --source=https://your-chocolatey-source.com
```

### Using Snap (Linux)
```bash
# Install from snap store (if published)
sudo snap install twig
```

## Docker (Alternative)

If you prefer using Docker:

```bash
# Pull the image
docker pull ghcr.io/yourusername/twig:latest

# Run with volume mount
docker run -v $(pwd):/workspace ghcr.io/yourusername/twig:latest

# Interactive mode
docker run -it -v $(pwd):/workspace ghcr.io/yourusername/twig:latest --serve
```

## Troubleshooting

### Permission Denied (Linux/macOS)
```bash
chmod +x twig
```

### Binary Not Found
- Ensure the binary is in your PATH
- Try running with full path: `./twig`

### Windows Issues
- If you get "Windows protected your PC" message, click "More info" → "Run anyway"
- Or run: `Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser`

### Version Check
```bash
twig --version
```

## File Structure

After installation, create this structure:
```
my-site/
├── content/          # Your markdown files
│   ├── index.md
│   └── about.md
├── template.html     # Custom template (optional)
└── public/           # Generated HTML (created by twig)
    ├── index.html
    └── about.html
```

## Next Steps

1. **Create your first site**: Follow the basic usage example above
2. **Customize the template**: Modify `template.html` for your design
3. **Add more content**: Create markdown files in the `content` directory
4. **Deploy**: Upload the `public` directory to any web hosting service

## Support

- **Issues**: [GitHub Issues](https://github.com/yourusername/twig/issues)
- **Documentation**: [README.md](https://github.com/yourusername/twig#readme)
- **Releases**: [GitHub Releases](https://github.com/yourusername/twig/releases) 