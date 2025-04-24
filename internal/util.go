package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// LogEntry represents a single log entry for file operations
type LogEntry struct {
	Timestamp   time.Time
	Action      string
	FilePath    string
	FileSize    int64
	IsDirectory bool
	Success     bool
	Error       string
}

// LogSummary represents a summary of cleaning operation
type LogSummary struct {
	StartTime           time.Time
	EndTime             time.Time
	TotalFilesScanned   int
	TotalBytesProcessed int64
	SuccessfulDeletes   int
	FailedDeletes       int
	DryRun              bool
}

// Logger handles logging operations
type Logger struct {
	entries []LogEntry
	summary LogSummary
}

// NewLogger creates a new Logger instance
func NewLogger(dryRun bool) *Logger {
	return &Logger{
		entries: make([]LogEntry, 0),
		summary: LogSummary{
			StartTime: time.Now(),
			DryRun:    dryRun,
		},
	}
}

// LogFileOperation logs a file operation
func (l *Logger) LogFileOperation(action string, path string, success bool, err error) {
	fileInfo, _ := os.Stat(path)
	var fileSize int64
	isDir := false

	if fileInfo != nil {
		fileSize = fileInfo.Size()
		isDir = fileInfo.IsDir()
	}

	// In dry-run mode, modify the action to indicate simulation
	if l.summary.DryRun && action == "DELETE" {
		action = "WOULD_DELETE"
	}

	entry := LogEntry{
		Timestamp:   time.Now(),
		Action:      action,
		FilePath:    path,
		FileSize:    fileSize,
		IsDirectory: isDir,
		Success:     success,
		Error:       "",
	}

	if err != nil {
		entry.Error = err.Error()
	}

	l.entries = append(l.entries, entry)

	// Update summary statistics
	if action == "DELETE" || action == "WOULD_DELETE" {
		if success {
			l.summary.SuccessfulDeletes++
			l.summary.TotalBytesProcessed += fileSize
		} else {
			l.summary.FailedDeletes++
		}
	}
	l.summary.TotalFilesScanned++
}

// FinalizeSummary completes the summary with end time
func (l *Logger) FinalizeSummary() LogSummary {
	l.summary.EndTime = time.Now()
	return l.summary
}

// PrintSummary prints a formatted summary of the operation
func (l *Logger) PrintSummary() {
	duration := l.summary.EndTime.Sub(l.summary.StartTime)

	fmt.Printf("\n📊 Operation Summary:\n")
	fmt.Printf("Start Time: %v\n", l.summary.StartTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("End Time: %v\n", l.summary.EndTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("Duration: %v\n", duration.Round(time.Millisecond))
	fmt.Printf("Total Files Scanned: %d\n", l.summary.TotalFilesScanned)
	fmt.Printf("Total Data Processed: %s\n", FormatBytes(l.summary.TotalBytesProcessed))

	if l.summary.DryRun {
		fmt.Printf("Files that would be deleted: %d\n", l.summary.SuccessfulDeletes)
	} else {
		fmt.Printf("Successfully Deleted: %d\n", l.summary.SuccessfulDeletes)
		fmt.Printf("Failed Deletions: %d\n", l.summary.FailedDeletes)
	}
}

// SaveLogToFile saves the log entries to a file
func (l *Logger) SaveLogToFile() error {
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return fmt.Errorf("failed to create log directory: %v", err)
	}

	// Add dry-run indication to log filename if in dry-run mode
	filename := "cleanup"
	if l.summary.DryRun {
		filename += "_dry_run"
	}
	filename += fmt.Sprintf("_%s.log", time.Now().Format("20060102_150405"))

	logFile := filepath.Join(logDir, filename)

	f, err := os.Create(logFile)
	if err != nil {
		return fmt.Errorf("failed to create log file: %v", err)
	}
	defer f.Close()

	// Write header with dry-run indication if applicable
	if l.summary.DryRun {
		fmt.Fprintf(f, "DRY RUN - Cleanup Operation Log - %s\n", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Fprintf(f, "NO FILES WERE ACTUALLY DELETED\n")
	} else {
		fmt.Fprintf(f, "Cleanup Operation Log - %s\n", time.Now().Format("2006-01-02 15:04:05"))
	}
	fmt.Fprintf(f, "----------------------------------------\n\n")

	// Write entries
	for _, entry := range l.entries {
		if entry.Action == "SCAN" && !entry.IsDirectory {
			continue // Skip logging scan operations for files in dry-run mode to reduce noise
		}
		fmt.Fprintf(f, "[%s] %s - %s\n",
			entry.Timestamp.Format("15:04:05"),
			entry.Action,
			entry.FilePath)
		if entry.Error != "" {
			fmt.Fprintf(f, "  Error: %s\n", entry.Error)
		}
	}

	// Write summary
	fmt.Fprintf(f, "\nSummary:\n")
	fmt.Fprintf(f, "Total Files Scanned: %d\n", l.summary.TotalFilesScanned)
	fmt.Fprintf(f, "Total Data Processed: %s\n", FormatBytes(l.summary.TotalBytesProcessed))
	fmt.Fprintf(f, "Successful Deletes: %d\n", l.summary.SuccessfulDeletes)
	fmt.Fprintf(f, "Failed Deletes: %d\n", l.summary.FailedDeletes)

	return nil
}

// FormatBytes converts bytes to human readable string
func FormatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// IsPathAccessible checks if a path is accessible
func IsPathAccessible(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// GetFileAge returns the age of a file in days
func GetFileAge(info os.FileInfo) float64 {
	return time.Since(info.ModTime()).Hours() / 24
}
