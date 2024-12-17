package output

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// WriteFilesToOutput writes the content of given files into a single output file
// in the format:
// // File: relative/path/to/file
// <file content>
func WriteFilesToOutput(outputFile, projectDir string, files []string) error {
	// Remove existing output file if exists
	if _, err := os.Stat(outputFile); err == nil {
		if err := os.Remove(outputFile); err != nil {
			return fmt.Errorf("failed to remove existing output file: %w", err)
		}
	}

	f, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer f.Close()

	for _, filePath := range files {
		relPath, err := filepath.Rel(projectDir, filePath)
		if err != nil {
			relPath = filePath // fallback to full path if relative fails
		}

		// Write header
		if _, err := io.WriteString(f, fmt.Sprintf("// File: %s\n", relPath)); err != nil {
			return fmt.Errorf("failed to write header: %w", err)
		}

		// Write file content
		file, err := os.Open(filePath)
		if err != nil {
			// skip this file if can't open
			continue
		}
		_, err = io.Copy(f, file)
		file.Close()
		if err != nil {
			return fmt.Errorf("failed to copy file content: %w", err)
		}

		// New line after file content
		if _, err := io.WriteString(f, "\n"); err != nil {
			return fmt.Errorf("failed to write newline: %w", err)
		}
	}

	return nil
}
