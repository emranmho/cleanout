package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type CleanResult struct {
	FilesChecked           int
	FilesMarkedForDeletion int
	FilesActuallyDeleted   int
}

func CleanDirectory(path string, days int, dryRun bool, verbose bool) (*CleanResult, error) {
	result := &CleanResult{}

	// Clean and normalize the path
	path = filepath.Clean(filepath.FromSlash(path))

	cutoff := time.Now().AddDate(0, 0, -days)

	err := filepath.Walk(path, func(f string, info os.FileInfo, err error) error {
		if err != nil {
			if verbose {
				fmt.Printf("⚠️  Skipping (access error): %s — %v\n", f, err)
			}
			return nil
		}

		result.FilesChecked++
		if info.IsDir() {
			if verbose {
				fmt.Printf("📂 Skipping directory: %s\n", f)
			}
			return nil
		}

		if info.ModTime().Before(cutoff) {
			result.FilesMarkedForDeletion++
			if dryRun {
				fmt.Printf("[DRY-RUN] Would delete: %s\n", f)
			} else {
				if verbose {
					fmt.Printf("🗑️  Deleting: %s\n", f)
				}
				if err := os.Remove(f); err != nil {
					fmt.Printf("❌ Failed to delete %s: %v\n", f, err)
				} else {
					result.FilesActuallyDeleted++
				}
			}
		} else if verbose {
			fmt.Printf("✅ Keeping: %s\n", f)
		}
		return nil
	})

	return result, err
}
