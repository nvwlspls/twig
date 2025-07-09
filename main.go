package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Version will be set by the linker during build
var version = "dev"

func main() {
	var sourceDir = flag.String("source", "content", "Source directory containing markdown files")
	var outputDir = flag.String("output", "public", "Output directory for generated HTML files")
	var templateFile = flag.String("template", "template.html", "HTML template file")
	var serve = flag.Bool("serve", false, "Serve the output directory locally after building")
	var port = flag.Int("port", 8080, "Port to serve the site on (used with --serve)")
	var showVersion = flag.Bool("version", false, "Show version information")
	var initProject = flag.Bool("init", false, "Initialize a new twig project in the current directory")
	flag.Parse()

	if *showVersion {
		fmt.Printf("twig version %s\n", version)
		os.Exit(0)
	}

	if *initProject {
		if err := initializeProject(); err != nil {
			log.Fatal("Failed to initialize project:", err)
		}
		fmt.Println("✅ Project initialized successfully!")
		fmt.Println("📝 Next steps:")
		fmt.Println("   1. Add markdown files to the 'content' directory")
		fmt.Println("   2. Run 'twig' to build your site")
		fmt.Println("   3. Run 'twig --serve' to preview locally")
		os.Exit(0)
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Fatal("Failed to create output directory:", err)
	}

	// Initialize the site generator
	generator := NewSiteGenerator(*sourceDir, *outputDir, *templateFile)

	// Build the site
	if err := generator.Build(); err != nil {
		log.Fatal("Failed to build site:", err)
	}

	fmt.Printf("Site built successfully! Output: %s\n", *outputDir)

	if *serve {
		fmt.Printf("Serving %s at http://localhost:%d ...\n", *outputDir, *port)
		httpServeDir(*outputDir, *port)
	}
}

// httpServeDir serves the given directory on the specified port
func httpServeDir(dir string, port int) {
	fs := http.FileServer(http.Dir(dir))
	addr := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(addr, fs); err != nil {
		log.Fatalf("Failed to serve %s: %v", dir, err)
	}
}

// initializeProject creates the basic directory structure and files for a new twig project
func initializeProject() error {
	// Create directories
	dirs := []string{"content", "public"}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}

	// Create sample content files
	sampleFiles := map[string]string{
		"content/index.md": `# Welcome to My Site

This is your homepage. Edit this file to customize your site.

## Getting Started

1. Add more markdown files to the 'content' directory
2. Run 'twig' to build your site
3. Run 'twig --serve' to preview locally

## Features

- **Markdown Support**: Write content in markdown
- **Auto-generated Index**: Your content automatically appears on the homepage
- **Customizable**: Modify the template to match your style
`,
		"content/about.md": `# About

This is the about page. You can add information about yourself, your project, or anything else.

## Contact

Feel free to reach out if you have any questions!
`,
		"content/blog/first-post.md": `# My First Blog Post

Welcome to my blog! This is my first post.

## Introduction

This is a sample blog post to get you started. You can:

- Write in **markdown** format
- Use *italics* and **bold** text
- Create [links](https://example.com)
- Add code blocks:

` + "```" + `go
func main() {
    fmt.Println("Hello, World!")
}
` + "```" + `

## Next Steps

1. Edit this file to add your own content
2. Add more posts to the 'content/blog' directory
3. Customize the template to match your style
`,
	}

	// Create the blog directory first
	if err := os.MkdirAll("content/blog", 0755); err != nil {
		return fmt.Errorf("failed to create blog directory: %v", err)
	}

	// Create sample files
	for filePath, content := range sampleFiles {
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to create file %s: %v", filePath, err)
		}
	}

	// Create a custom template if it doesn't exist
	if _, err := os.Stat("template.html"); os.IsNotExist(err) {
		customTemplate := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            line-height: 1.6;
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
            background-color: #f8f9fa;
        }
        .container {
            background: white;
            padding: 40px;
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
        }
        h1 { color: #2c3e50; border-bottom: 3px solid #3498db; padding-bottom: 10px; }
        h2 { color: #34495e; margin-top: 30px; }
        h3 { color: #7f8c8d; }
        a { color: #3498db; text-decoration: none; }
        a:hover { text-decoration: underline; }
        code { background: #f1f2f6; padding: 2px 6px; border-radius: 3px; }
        pre { background: #2c3e50; color: #ecf0f1; padding: 20px; border-radius: 5px; overflow-x: auto; }
        pre code { background: none; padding: 0; }
        .nav { margin-bottom: 30px; padding: 10px 0; border-bottom: 1px solid #eee; }
        .nav a { margin-right: 20px; }
        .footer { margin-top: 50px; padding-top: 20px; border-top: 1px solid #eee; color: #7f8c8d; }
    </style>
</head>
<body>
    <div class="container">
        <div class="nav">
            <a href="index.html">Home</a>
            <a href="about.html">About</a>
            <a href="blog/">Blog</a>
        </div>
        
        <div class="content">
            {{.Content}}
        </div>
        
        <div class="footer">
            <p>Generated with <a href="https://github.com/yourusername/twig">twig</a></p>
        </div>
    </div>
</body>
</html>`
		
		if err := os.WriteFile("template.html", []byte(customTemplate), 0644); err != nil {
			return fmt.Errorf("failed to create template.html: %v", err)
		}
	}

	// Create a README file
	readmeContent := `# My Twig Site

This is a static site generated with [twig](https://github.com/yourusername/twig).

## Structure

- 'content/' - Your markdown files
- 'public/' - Generated HTML files
- 'template.html' - HTML template

## Usage

1. Add markdown files to the 'content' directory
2. Run 'twig' to build your site
3. Run 'twig --serve' to preview locally
4. Deploy the 'public' directory to any web hosting service

## Customization

- Edit 'template.html' to change the site design
- Add CSS to the template for custom styling
- Organize content in subdirectories within 'content/'

Happy building! 🚀
`

	if err := os.WriteFile("README.md", []byte(readmeContent), 0644); err != nil {
		return fmt.Errorf("failed to create README.md: %v", err)
	}

	return nil
}
