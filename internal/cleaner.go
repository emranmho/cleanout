package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// CleanResult contains the statistics of the cleaning operation
type CleanResult struct {
	FilesChecked           int
	FilesMarkedForDeletion int
	FilesActuallyDeleted   int
}

// CleanDirectory scans a directory and deletes/marks files older than specified days
func CleanDirectory(path string, days int, dryRun bool, verbose bool) (*CleanResult, error) {
	result := &CleanResult{}

	// Clean and normalize the path
	path = filepath.Clean(filepath.FromSlash(path))

	cutoff := time.Now().AddDate(0, 0, -days)

	err := filepath.Walk(path, func(f string, info os.FileInfo, err error) error {
		if err != nil {
			if verbose {
				fmt.Printf("Skipping %s: %v\n", f, err)
			}
			return nil
		}

		result.FilesChecked++
		if info.IsDir() {
			if verbose {
				fmt.Printf("Skipping directory: %s\n", f)
			}
			return nil
		}

		if info.ModTime().Before(cutoff) {
			result.FilesMarkedForDeletion++
			if dryRun {
				fmt.Println("[DRY-RUN] Would delete:", f)
			} else {
				if verbose {
					fmt.Println("Deleting:", f)
				}
				if err := os.Remove(f); err != nil {
					fmt.Printf("Failed to delete %s: %v\n", f, err)
				} else {
					result.FilesActuallyDeleted++
				}
			}
		} else if verbose {
			fmt.Println("Keeping:", f)
		}
		return nil
	})

	return result, err
}
