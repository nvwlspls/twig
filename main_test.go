package main

import (
	"os"
	"testing"
)

func TestMainFunction(t *testing.T) {
	// Test that the program can be built and run
	// This is a basic smoke test
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}
}

func TestFileOperations(t *testing.T) {
	// Test that we can create directories
	testDir := "test_output"
	err := os.MkdirAll(testDir, 0755)
	if err != nil {
		t.Errorf("Failed to create test directory: %v", err)
	}
	defer os.RemoveAll(testDir)
}

func TestSiteGeneratorCreation(t *testing.T) {
	// Test that we can create a site generator
	generator := NewSiteGenerator("content", "test_output", "template.html")
	if generator == nil {
		t.Error("Failed to create site generator")
	}
} 