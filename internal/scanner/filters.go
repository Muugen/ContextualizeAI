package scanner

import (
	"os"
	"strings"
)

// IsIgnored checks if the given path should be ignored based on
// directory naming conventions and ignore patterns from config.
// This function:
// - Ignores directories that start with '.'
// - Ignores any path containing 'node_modules'
// - Applies ignore patterns from config
func IsIgnored(path string, info os.FileInfo, ignorePatterns []string) bool {
	// Normalize path and convert to lower for case-insensitive matching
	lowerPath := strings.ToLower(path)

	// Ignore directories starting with '.'
	if info.IsDir() && strings.HasPrefix(info.Name(), ".") {
		return true
	}

	// Ignore anything within 'node_modules'
	if strings.Contains(lowerPath, "node_modules") {
		return true
	}

	// Apply ignore patterns (from config)
	for _, pat := range ignorePatterns {
		if matchPattern(lowerPath, pat) {
			return true
		}
	}

	return false
}

// matchPattern is a basic pattern matcher supporting:
// - Prefix/suffix wildcards like "*something" or "something*"
// - Both ends wildcard "*something*"
// - Otherwise checks if the pattern appears as a substring.
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
