package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                        "app",
	Short:                      "A sample CLI program",
	Long:                       "A sample CLI program that demonstrates intelligent command suggestions",
	SuggestionsMinimumDistance: 2,
	DisableSuggestions:         false,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
