package commands

import (
	"terminal-todos/internal/daemon"

	"github.com/spf13/cobra"
)

var debug bool

var serveCmd = &cobra.Command{
	Use: "/serve",
	Short: "serve - starts service in control mode",
	Run: func (cmd *cobra.Command, argv []string) {
		service := daemon.New()
		service.Start(debug)
	},
}

func init() {
	serveCmd.Flags().BoolVarP(&debug, "debug", "d", false, "enable debug")
	rootCmd.AddCommand(serveCmd)
}