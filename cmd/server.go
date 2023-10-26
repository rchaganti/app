package cmd

import (
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "This is the top-level server command",
	Long:  "This is the top-level server command",
	Args:  cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
