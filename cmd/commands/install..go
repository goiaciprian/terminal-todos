package commands

import (
	"terminal-todos/internal/config"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install - creates the required files",
	Run: func(cmd *cobra.Command, argv []string) {
		config.FirstTimeSetup()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
