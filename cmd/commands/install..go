package commands

import (
	"terminal-todos/internal/config"

	"github.com/spf13/cobra"
)

var migrationsFolder string

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install - creates the required files",
	Run: func(cmd *cobra.Command, argv []string) {
		config.FirstTimeSetup(migrationsFolder)
	},
}

func init() {
	installCmd.Flags().StringVarP(&migrationsFolder, "migrationFolder", "m", "", "migration folder for db")
	installCmd.MarkFlagRequired("migrationFolder")
	rootCmd.AddCommand(installCmd)
}
