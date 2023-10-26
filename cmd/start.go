package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Stars a server",
	Long:  "Starts a server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting a server!") // Do something
	},
}

func init() {
	serverCmd.AddCommand(startCmd)
}
