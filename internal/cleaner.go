package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// CleanResult holds the result of the cleaning operation
type CleanResult struct {
	FilesChecked           int
	FilesMarkedForDeletion int
	FilesActuallyDeleted   int
}

// CleanDirectory scans a directory and deletes/marks files older than specified days
func CleanDirectory(path string, days int, dryRun bool, verbose bool) (*CleanResult, error) {
	result := &CleanResult{}
	logger := NewLogger(dryRun)

	// Clean and normalize the path
	path = filepath.Clean(filepath.FromSlash(path))

	if !IsPathAccessible(path) {
		return result, fmt.Errorf("path is not accessible: %s", path)
	}

	cutoff := time.Now().AddDate(0, 0, -days)

	err := filepath.Walk(path, func(f string, info os.FileInfo, err error) error {
		if err != nil {
			logger.LogFileOperation("SCAN", f, false, err)
			if verbose {
				fmt.Printf("⚠️ Skipping %s: %v\n", f, err)
			}
			return nil
		}

		result.FilesChecked++
		logger.LogFileOperation("SCAN", f, true, nil)

		if info.IsDir() {
			if verbose {
				fmt.Printf("📂 Skipping directory: %s\n", f)
			}
			return nil
		}

		if info.ModTime().Before(cutoff) {
			result.FilesMarkedForDeletion++
			if dryRun {
				fmt.Println("[DRY-RUN] Would delete:", f)
				logger.LogFileOperation("DELETE", f, true, nil)
			} else {
				if verbose {
					fmt.Println("🗑️ Deleting:", f)
				}
				err := os.Remove(f)
				logger.LogFileOperation("DELETE", f, err == nil, err)
				if err != nil {
					fmt.Printf("❌ Failed to delete %s: %v\n", f, err)
				} else {
					result.FilesActuallyDeleted++
				}
			}
		} else if verbose {
			fmt.Println("✅ Keeping:", f)
		}
		return nil
	})

	// Finalize and save logs
	logger.FinalizeSummary()
	logger.PrintSummary()
	if err := logger.SaveLogToFile(); err != nil && verbose {
		fmt.Printf("Warning: Failed to save log file: %v\n", err)
	}

	return result, err
}
