package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var logDays, totalFiledDeleted int

func init() {
	rootCmd.AddCommand(autoCleanLogCmd)

	autoCleanLogCmd.Flags().IntVar(&logDays, "days", 7, "Delete log files older than this many days")
}

var autoCleanLogCmd = &cobra.Command{
	Use:   "clean-logs",
	Short: "Delete old log files from the logs directory",
	Long:  "Manually delete log files older than a given number of days from the logs directory.",
	Example: `
# Delete logs older than 7 days (default)
  cleanout clean-logs

  # Delete logs older than 14 days
  cleanout clean-logs --days 14
	`,
	Run: func(cmd *cobra.Command, args []string) {

		logsDir := "logs"
		cutOff := time.Now().AddDate(0, 0, -logDays)

		fmt.Println("🧹 Cleaning log files older than", logDays, " days in ", logsDir)

		// Check if the logs directory exists
		if _, err := os.Stat(logsDir); os.IsNotExist(err) {
			cmd.Println("Logs directory does not exist:", logsDir)
			return
		}

		// Call the function to clean old log files
		err := filepath.Walk(logsDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("⚠️ Failed to access %s: %v\n", path, err)
				return nil
			}
			// Check if the file is older than the cut-off date

			if info.ModTime().Before(cutOff) {
				if !info.IsDir() {
					totalFiledDeleted++
					// Delete the file
					if err := os.Remove(path); err != nil {
						cmd.Println("❌ Error deleting file:", path, err)
					} else {
						cmd.Println("🗑️ Deleted file:", path)
					}
				}
			}
			return nil
		})

		if err != nil {
			fmt.Printf("❌ Error during log cleanup: %v\n", err)
		} else {
			fmt.Println("✅ Log cleanup complete.")
			fmt.Printf("Total files deleted: %d\n", totalFiledDeleted)
		}
	},
}
