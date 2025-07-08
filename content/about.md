# About Twig

**Twig** is a minimal static site generator written in Go. It's designed to be simple, fast, and easy to use.

## Why Twig?

- **Minimal**: No complex configuration or dependencies
- **Fast**: Written in Go for excellent performance
- **Simple**: Just markdown files and a template
- **Flexible**: Customizable templates and output

## How It Works

1. **Input**: Markdown files in a source directory
2. **Processing**: Convert markdown to HTML using Go's markdown parser
3. **Templating**: Apply HTML template with metadata
4. **Output**: Generate static HTML files

## Getting Started

1. Create a `content` directory with your markdown files
2. Optionally customize the `template.html`
3. Run `go run .` to build your site
4. View the generated files in the `public` directory

## Customization

You can customize:
- **Template**: Modify `template.html` for different styling
- **Source directory**: Use `--source` flag
- **Output directory**: Use `--output` flag
- **Template file**: Use `--template` flag

## Future Enhancements

- [ ] Front matter support (YAML/TOML)
- [ ] Custom CSS/JS injection
- [ ] RSS feed generation
- [ ] Tag/category support
- [ ] Pagination
