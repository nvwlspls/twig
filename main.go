package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var sourceDir = flag.String("source", "content", "Source directory containing markdown files")
	var outputDir = flag.String("output", "public", "Output directory for generated HTML files")
	var templateFile = flag.String("template", "template.html", "HTML template file")
	var serve = flag.Bool("serve", false, "Serve the output directory locally after building")
	var port = flag.Int("port", 8080, "Port to serve the site on (used with --serve)")
	flag.Parse()

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
