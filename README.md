# Twig - Minimal Static Site Generator

A simple, fast static site generator written in Go that converts markdown files to HTML.

## Features

- **Markdown Support**: Convert `.md` files to HTML with full markdown syntax
- **Clean Templates**: Beautiful, responsive HTML templates
- **Auto-generated Index**: Automatic index page with all your content
- **Simple CLI**: Easy-to-use command line interface
- **Customizable**: Flexible template and output options

## Installation

### Quick Install

**Linux/macOS:**
```bash
curl -fsSL https://raw.githubusercontent.com/yourusername/twig/main/scripts/install.sh | bash
```

**Windows (PowerShell):**
```powershell
Invoke-Expression (Invoke-WebRequest -Uri "https://raw.githubusercontent.com/yourusername/twig/main/scripts/install.ps1" -UseBasicParsing).Content
```

### Manual Installation

1. **Download from [GitHub Releases](https://github.com/yourusername/twig/releases)**
2. **Extract the archive for your platform**
3. **Add to PATH** (see [INSTALL.md](INSTALL.md) for detailed instructions)

### Package Managers

**Homebrew (macOS):**
```bash
brew install yourusername/twig/twig
```

**Chocolatey (Windows):**
```powershell
choco install twig
```

**Go Install (if Go is installed):**
```bash
go install github.com/yourusername/twig@latest
```

### From Source

1. Clone this repository:
```bash
git clone <repository-url>
cd twig
```

2. Install dependencies:
```bash
go mod tidy
```

## Usage

### Basic Usage

1. Create markdown files in a `content` directory
2. Run the generator:
```bash
go run .
```

This will:
- Read markdown files from `content/`
- Generate HTML files in `public/`
- Create an index page at `public/index.html`

### Command Line Options

```bash
go run . [options]

Options:
  --source string     Source directory containing markdown files (default "content")
  --output string     Output directory for generated HTML files (default "public")
  --template string   HTML template file (default "template.html")
```

### Examples

```bash
# Use custom directories
go run . --source posts --output site

# Use custom template
go run . --template custom.html

# Full custom configuration
go run . --source blog --output docs --template blog.html
```

## File Structure

```
twig/
├── main.go              # Main entry point
├── generator.go          # Site generator logic
├── template.html         # Default HTML template
├── go.mod               # Go module file
├── content/             # Source markdown files
│   ├── hello-world.md
│   └── about.md
└── public/              # Generated HTML files (created after build)
    ├── index.html
    ├── hello-world.html
    └── about.html
```

## Markdown Features

The generator supports standard markdown syntax:

- **Headers**: `# H1`, `## H2`, etc.
- **Bold/Italic**: `**bold**`, `*italic*`
- **Lists**: Ordered and unordered lists
- **Links**: `[text](url)`
- **Code**: Inline `code` and code blocks
- **Blockquotes**: `> quote`
- **Tables**: Standard markdown tables

## Customization

### Template

The `template.html` file contains the HTML template used for all pages. It includes:

- Responsive CSS styling
- Navigation back to index
- Page metadata display
- Clean typography

You can modify this file to change the appearance of your site.

### Styling

The template includes built-in CSS for:
- Responsive design
- Syntax highlighting for code blocks
- Clean typography
- Modern color scheme

## Building

To build a standalone executable:

```bash
go build -o twig .
```

Then use it like:
```bash
./twig --source content --output public
```

### Version Information

Check the version of your twig binary:

```bash
./twig --version
```

## Release Management

This project uses automated versioning and releases via GitHub Actions.

### Automatic Versioning

The project supports automatic version bumping based on conventional commits:

- **Major version** (`1.0.0` → `2.0.0`): Breaking changes (`BREAKING CHANGE:` or `!:`)
- **Minor version** (`1.0.0` → `1.1.0`): New features (`feat:`)
- **Patch version** (`1.0.0` → `1.0.1`): Bug fixes and other changes

### Manual Version Bumping

Use the local script to bump versions:

```bash
# Bump patch version (default)
./scripts/bump-version.sh

# Bump minor version
./scripts/bump-version.sh minor

# Bump major version
./scripts/bump-version.sh major
```

### Creating Releases

#### Option 1: Manual Release
1. Bump version locally: `./scripts/bump-version.sh minor`
2. Commit changes: `git add VERSION && git commit -m "chore: bump version to 1.1.0"`
3. Create tag: `git tag -a "v1.1.0" -m "Release v1.1.0"`
4. Push: `git push origin main --tags`

#### Option 2: GitHub Actions Manual Trigger
1. Go to Actions → "Version Bump" workflow
2. Click "Run workflow"
3. Select bump type and run

#### Option 3: Automatic on Push to Main
- Push conventional commits to main branch
- GitHub Actions will automatically bump version and create releases

### Release Artifacts

Each release includes binaries for:
- **Windows**: amd64, 386
- **Linux**: amd64, 386, arm64  
- **macOS**: amd64, arm64

Download from: [GitHub Releases](https://github.com/yourusername/twig/releases)

## Example Output

After running the generator, you'll get:

- `public/index.html` - Index page listing all content
- `public/hello-world.html` - Generated from `content/hello-world.md`
- `public/about.html` - Generated from `content/about.md`

## Development

### Adding Features

The generator is modular and easy to extend:

- **New markdown features**: Modify the markdown parser options in `generator.go`
- **Custom metadata**: Extend the `Page` struct and parsing logic
- **Additional templates**: Add template functions and variables

### Dependencies

- `github.com/gomarkdown/markdown` - Markdown parsing and rendering

## License

MIT License - feel free to use and modify as needed.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test with `go run .`
5. Submit a pull request
