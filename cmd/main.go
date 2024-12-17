package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"ContextualizeAI/internal/config"
	"ContextualizeAI/internal/output"
	"ContextualizeAI/internal/scanner"
	"ContextualizeAI/internal/ui"
)

func main() {
	// Add a flag for project directory
	projectDirFlag := flag.String("project-dir", "", "Path to the project directory to scan")
	flag.Parse()

	var projectDir string
	if *projectDirFlag != "" {
		projectDir = *projectDirFlag
	} else {
		// fallback to current directory if not provided
		dir, err := os.Getwd()
		if err != nil {
			log.Fatalf("Unable to get current directory: %v\n", err)
		}
		projectDir = dir
	}

	// Load configuration
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
	}

	// Find top-level directories in projectDir
	dirs, err := topLevelDirectories(projectDir)
	if err != nil {
		log.Fatalf("Error listing directories: %v\n", err)
	}

	// Interactive selection of directories
	selectedDirs, includeRootFiles := ui.SelectDirectories(dirs)

	// Gather files based on selected directories
	var allFiles []string
	if includeRootFiles {
		rootFiles, err := scanner.ListFiles(projectDir, cfg.IgnorePatterns)
		if err != nil {
			log.Printf("Warning: Error listing root files: %v\n", err)
		} else {
			allFiles = append(allFiles, rootFiles...)
		}
	}

	for _, d := range selectedDirs {
		fullDir := filepath.Join(projectDir, d)
		files, err := scanner.ListFiles(fullDir, cfg.IgnorePatterns)
		if err != nil {
			log.Printf("Warning: Error listing files in %s: %v\n", d, err)
			continue
		}
		allFiles = append(allFiles, files...)
	}

	// Write output
	if err := output.WriteFilesToOutput(cfg.OutputFile, projectDir, allFiles); err != nil {
		log.Fatalf("Error writing output: %v\n", err)
	}

	fmt.Println("MVP execution completed. Output file:", cfg.OutputFile)
}

// topLevelDirectories returns top-level subdirectories from the given directory
func topLevelDirectories(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var dirs []string
	for _, e := range entries {
		if e.IsDir() {
			dirs = append(dirs, e.Name())
		}
	}
	return dirs, nil
}