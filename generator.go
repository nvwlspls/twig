package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
)

// Page represents a single page with its metadata and content
type Page struct {
	Title       string
	Content     template.HTML
	Date        time.Time
	URL         string
	Description string
}

// SiteGenerator handles the site generation process
type SiteGenerator struct {
	sourceDir    string
	outputDir    string
	templateFile string
	pages        []Page
}

// NewSiteGenerator creates a new site generator instance
func NewSiteGenerator(sourceDir, outputDir, templateFile string) *SiteGenerator {
	return &SiteGenerator{
		sourceDir:    sourceDir,
		outputDir:    outputDir,
		templateFile: templateFile,
		pages:        []Page{},
	}
}

// Build processes all markdown files and generates the site
func (sg *SiteGenerator) Build() error {
	// Find all markdown files
	files, err := sg.findMarkdownFiles()
	if err != nil {
		return fmt.Errorf("failed to find markdown files: %w", err)
	}

	// Process each markdown file
	for _, file := range files {
		if err := sg.processFile(file); err != nil {
			log.Printf("Warning: failed to process %s: %v", file, err)
			continue
		}
	}

	// Sort pages by date for consistent ordering
	// This ensures a predictable order in the sidebar
	sort.Slice(sg.pages, func(i, j int) bool {
		return sg.pages[i].Date.After(sg.pages[j].Date)
	})

	// Regenerate all HTML files with the complete sorted page list
	for _, page := range sg.pages {
		if err := sg.regenerateHTMLFile(page); err != nil {
			log.Printf("Warning: failed to regenerate %s: %v", page.URL, err)
		}
	}

	// Generate index page
	if err := sg.generateIndex(); err != nil {
		return fmt.Errorf("failed to generate index: %w", err)
	}

	return nil
}

// findMarkdownFiles recursively finds all .md files in the source directory
func (sg *SiteGenerator) findMarkdownFiles() ([]string, error) {
	var files []string

	err := filepath.Walk(sg.sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".md") {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

// processFile converts a single markdown file to HTML
func (sg *SiteGenerator) processFile(filePath string) error {
	// Read markdown content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Parse markdown to HTML
	html := markdown.ToHTML(content, nil, nil)

	// Extract metadata and content
	page := sg.extractPageInfo(filePath, content, html)

	// Add the page to pages collection before generating HTML
	sg.pages = append(sg.pages, page)

	// Generate HTML file
	if err := sg.generateHTMLFile(page); err != nil {
		return err
	}

	return nil
}

// extractPageInfo extracts metadata and content from markdown
func (sg *SiteGenerator) extractPageInfo(filePath string, markdownContent, htmlContent []byte) Page {
	// Default title from filename
	title := strings.TrimSuffix(filepath.Base(filePath), ".md")
	title = strings.ReplaceAll(title, "-", " ")
	title = strings.ReplaceAll(title, "_", " ")

	// Generate URL
	relPath, _ := filepath.Rel(sg.sourceDir, filePath)
	url := strings.TrimSuffix(relPath, ".md") + ".html"
	if url == "index.html" {
		url = "index.html"
	}

	// Try to extract title from first heading
	lines := strings.Split(string(markdownContent), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "# ") {
			title = strings.TrimPrefix(line, "# ")
			break
		}
	}

	return Page{
		Title:   title,
		Content: template.HTML(htmlContent),
		Date:    time.Now(),
		URL:     url,
	}
}

// generateHTMLFile creates an HTML file for a page
func (sg *SiteGenerator) generateHTMLFile(page Page) error {
	// Create output file path
	outputPath := filepath.Join(sg.outputDir, page.URL)
	outputDir := filepath.Dir(outputPath)

	// Make sure the output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	// Read template
	tmpl, err := template.ParseFiles(sg.templateFile)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Create output file
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create template data with current page and all pages
	templateData := struct {
		Page
		AllPages []Page
	}{
		Page:     page,
		AllPages: sg.pages,
	}

	// Execute template
	return tmpl.Execute(file, templateData)
}

// regenerateHTMLFile is an alias for generateHTMLFile to maintain compatibility with the updated Build function
func (sg *SiteGenerator) regenerateHTMLFile(page Page) error {
	return sg.generateHTMLFile(page)
}

// generateIndex creates an index page listing all pages
func (sg *SiteGenerator) generateIndex() error {
	// Create index template
	indexTmpl := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Site Index</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            line-height: 1.6;
            color: #333;
            margin: 0;
            padding: 0;
            background-color: #fafafa;
        }

        .layout {
            display: flex;
            min-height: 100vh;
        }

        .sidebar {
            width: 250px;
            background: white;
            border-right: 1px solid #eee;
            padding: 20px;
            position: fixed;
            height: 100vh;
            overflow-y: auto;
        }

        .main-content {
            flex: 1;
            margin-left: 250px;
            padding: 20px;
        }

        .container {
            background: white;
            padding: 40px;
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
            max-width: 800px;
            margin: 0 auto;
        }

        .sidebar h3 {
            color: #2c3e50;
            margin-bottom: 15px;
            font-size: 1.2em;
        }

        .sidebar ul {
            list-style: none;
            padding: 0;
            margin: 0;
        }

        .sidebar li {
            margin-bottom: 8px;
        }

        .sidebar a {
            color: #666;
            text-decoration: none;
            display: block;
            padding: 8px 12px;
            border-radius: 4px;
            transition: background-color 0.2s;
        }

        .sidebar a:hover {
            background-color: #f8f9fa;
            color: #3498db;
        }

        .sidebar a.active {
            background-color: #3498db;
            color: white;
        }

        .page-list { list-style: none; padding: 0; }
        .page-list li { margin: 10px 0; padding: 10px; border: 1px solid #ddd; border-radius: 5px; }
        .page-list a { text-decoration: none; color: #333; font-weight: bold; }
        .page-list a:hover { color: #007acc; }
    </style>
</head>
<body>
    <div class="layout">
        <div class="sidebar">
            <h3>Posts</h3>
            <ul>
                {{range .}}
                <li>
                    <a href="{{.URL}}" {{if eq .URL "index.html"}}class="active"{{end}}>
                        {{.Title}}
                    </a>
                </li>
                {{end}}
            </ul>
        </div>

        <div class="main-content">
            <div class="container">
                <h1>Site Index</h1>
                <ul class="page-list">
                    {{range .}}
                    <li>
                        <a href="{{.URL}}">{{.Title}}</a>
                        <br><small>{{.Date.Format "2006-01-02"}}</small>
                    </li>
                    {{end}}
                </ul>
            </div>
        </div>
    </div>
</body>
</html>`

	tmpl, err := template.New("index").Parse(indexTmpl)
	if err != nil {
		return err
	}

	// Create index file
	indexPath := filepath.Join(sg.outputDir, "index.html")
	file, err := os.Create(indexPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, sg.pages)
}
