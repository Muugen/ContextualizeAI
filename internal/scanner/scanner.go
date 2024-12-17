package scanner

import (
	"os"
	"path/filepath"
)

func ListFiles(dir string, ignorePatterns []string) ([]string, error) {
	var result []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Skip this file/dir on error
			return nil
		}

		if IsIgnored(path, info, ignorePatterns) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if !info.IsDir() {
			result = append(result, path)
		}
		return nil
	})

	return result, err
}
