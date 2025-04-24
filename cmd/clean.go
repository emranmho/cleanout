package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/emranmho/cleanout/internal"
	"github.com/spf13/cobra"
)

var (
	path    string // Path to scan for temp/cache files
	days    int    // Age in days to consider a file old
	dryRun  bool   // Flag for dry run mode
	verbose bool   // Flag for verbose output
)

func init() {
	rootCmd.AddCommand(cleanCmd)

	cleanCmd.Flags().StringVar(&path, "path", os.TempDir(), "Directory to scan")
	cleanCmd.Flags().IntVar(&days, "days", 7, "Delete files older than N days")
	cleanCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Preview deletions only")
	cleanCmd.Flags().BoolVar(&verbose, "verbose", false, "Show detailed logs")
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean old files from a directory",
	Long: `Scans a directory and deletes files older than N days.
Supports dry-run mode to preview deletions.`,
	Example: `
# 🧹 Default Cleanup
# If you just run 'clean' with no flags, it will:
# - Use the system temp directory (like C:\Users\You\AppData\Local\Temp or /tmp)
# - Delete files older than 7 days
cleanout clean

# 🔍 Dry-run Mode
# Preview what files would be deleted without actually deleting them
cleanout clean --dry-run

# 🕒 Customize Age of Files
# Delete files older than 30 days
cleanout clean --days 30

# 📁 Custom Path
# Clean old files in a specific folder
cleanout clean --path "C:\MyCache" --days 10

# 🔊 Verbose Mode
# Show detailed log while deleting
cleanout clean --path /tmp --days 5 --verbose

# 💡 Combine flags
# Full command: clean a custom folder, show details, preview only
cleanout clean --path ~/Downloads --days 14 --dry-run --verbose
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Clean and normalize the path
		path = filepath.Clean(filepath.FromSlash(path))

		// Check if directory exists
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Printf("Error: Directory does not exist: %s\n", path)
			if strings.Contains(path, " ") {
				fmt.Println("Hint: If your path contains spaces, wrap it in double quotes.")
			}
			return
		}

		if verbose {
			fmt.Println("🔍 Scanning:", path)
			fmt.Printf("🕒 Looking for files older than %d days...\n", days)
		}

		result, err := internal.CleanDirectory(path, days, dryRun, verbose)
		if err != nil {
			fmt.Printf("Error during cleanup: %v\n", err)
			return
		}

		// Print detailed summary
		fmt.Printf("\n📊 Summary:\n")
		fmt.Printf("✓ Files checked: %d\n", result.FilesChecked)
		if dryRun {
			fmt.Printf("🗑️ Files marked for deletion: %d\n", result.FilesMarkedForDeletion)
			fmt.Println("ℹ️ No files were actually deleted (dry-run mode)")
		} else {
			fmt.Printf("🗑️ Files deleted successfully: %d\n", result.FilesActuallyDeleted)
			if result.FilesActuallyDeleted < result.FilesMarkedForDeletion {
				fmt.Printf("⚠️ Failed to delete %d files\n", result.FilesMarkedForDeletion-result.FilesActuallyDeleted)
			}
		}
	},
}
