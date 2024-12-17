package scanner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ListFiles scans a directory recursively and returns a list of file paths
// that do not match the ignore patterns. This is a basic version without concurrency.
func ListFiles(dir string, ignorePatterns []string) ([]string, error) {
	var result []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Skip this file/dir on error
			return nil
		}

		if info.IsDir() && path != dir {
			// Check if directory matches ignore patterns
			if shouldIgnore(path, ignorePatterns) {
				return filepath.SkipDir
			}
			return nil
		}

		// Check files
		if !info.IsDir() {
			if !shouldIgnore(path, ignorePatterns) {
				result = append(result, path)
			}
		}
		return nil
	})

	return result, err
}

func shouldIgnore(path string, patterns []string) bool {
	lowerPath := strings.ToLower(path)
	for _, pat := range patterns {
		if matchPattern(lowerPath, pat) {
			return true
		}
	}
	return false
}

// matchPattern is a simplified matching function:
// - If pattern starts with '*', match suffix
// - If pattern ends with '*', match prefix
// - Otherwise, exact substring match.
// This is basic and can be replaced with more advanced regex/glob logic later.
func matchPattern(path, pattern string) bool {
	pattern = strings.ToLower(pattern)
	if strings.HasPrefix(pattern, "*") && strings.HasSuffix(pattern, "*") {
		// *something*
		mid := strings.Trim(pattern, "*")
		return strings.Contains(path, mid)
	} else if strings.HasPrefix(pattern, "*") {
		// *something
		suffix := strings.TrimPrefix(pattern, "*")
		return strings.HasSuffix(path, suffix)
	} else if strings.HasSuffix(pattern, "*") {
		// something*
		prefix := strings.TrimSuffix(pattern, "*")
		return strings.HasPrefix(path, prefix)
	}
	// exact substring match
	return strings.Contains(path, pattern)
}
