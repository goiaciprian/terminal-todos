package commands

import (
	"fmt"
	"os"
	"strings"
	"terminal-todos/internal/todo"

	"github.com/spf13/cobra"
)

var description string

var createCmd = &cobra.Command{
	Use:   "create <title>",
	Short: "create - creates a new todo",
	Run: func(cmd *cobra.Command, args []string) {
		title := strings.Join(args, " ")
		todos, err := todo.Create(title, description)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			todo.Display(&todos)
		}

	},
}

func init() {
	createCmd.Flags().StringVarP(&description, "description", "d", "", "add description for todo")
	rootCmd.AddCommand(createCmd)
}
