package cmd

import (
	"github.com/spf13/cobra"
)

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Manage GitHub Apps",
}

func init() {
	rootCmd.AddCommand(appCmd)
}
