package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "cleanout",
	Short: "A CLI tool to clean temp/cache files older than a specified age",
	Long: `Cleanout is a CLI utility for scanning and removing old temp/cache files.
It supports dry-run mode, custom paths, and age threshold.`,
}

func Execute() {
	// Execute the root command
	cobra.CheckErr(rootCmd.Execute())
}
