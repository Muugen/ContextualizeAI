package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// SelectDirectories presents a menu for selecting directories and optionally including root files.
func SelectDirectories(allDirs []string) ([]string, bool) {
	reader := bufio.NewReader(os.Stdin)
	selected := make(map[string]bool)
	includeRootFiles := false

	for {
		fmt.Println("\nSelect directories to process:")
		fmt.Println("  1) All")
		for i, d := range allDirs {
			prefix := "   "
			if selected[d] {
				prefix = "  *"
			}
			fmt.Printf("%s%d) %s\n", prefix, i+2, d)
		}
		fmt.Printf("  %d) Root files (currently: %v)\n", len(allDirs)+2, includeRootFiles)
		fmt.Printf("  %d) Finish selection\n", len(allDirs)+3)
		fmt.Print("Enter your choice: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid choice. Please enter a number.")
			continue
		}

		switch {
		case choice == 1:
			// Select all directories
			for _, d := range allDirs {
				selected[d] = true
			}
			return mapKeys(selected), includeRootFiles
		case choice >= 2 && choice <= len(allDirs)+1:
			dir := allDirs[choice-2]
			if selected[dir] {
				delete(selected, dir)
				fmt.Println("Removed:", dir)
			} else {
				selected[dir] = true
				fmt.Println("Added:", dir)
			}
		case choice == len(allDirs)+2:
			includeRootFiles = !includeRootFiles
			fmt.Println("Toggled root files to:", includeRootFiles)
		case choice == len(allDirs)+3:
			if len(selected) == 0 && !includeRootFiles {
				fmt.Println("No directories or root files selected. Please choose at least one.")
			} else {
				return mapKeys(selected), includeRootFiles
			}
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func mapKeys(m map[string]bool) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
