package commands

import (
	"terminal-todos/internal/database"

	"github.com/spf13/cobra"
)


var installCmd = &cobra.Command{
	Use: "install",
	Short: "install - creates the required files",
	Run: func (cmd *cobra.Command, argv []string) {
		database.FirstTimeSetup()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}