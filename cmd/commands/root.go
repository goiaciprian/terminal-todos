package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version = "dev-build"

var rootCmd = &cobra.Command{
	Use:   "todos",
	Short: "todos - local cli todos app",
	Long:  "todos is suppose to help you add todos from you're terminal",
	Version: Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing your command '%s'", err)
		os.Exit(1)
	}
}
