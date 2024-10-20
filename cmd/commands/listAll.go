package commands

import (
	"fmt"
	"os"
	"terminal-todos/internal/todo"

	"github.com/spf13/cobra"
)


var listaAllCMd = &cobra.Command{
	Use: "list",
	Short: "list - display all todos",
	Run: func (cmd *cobra.Command, argv []string) {
		todos, err := todo.GetAll()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting the todos %s", err)
		}

		todo.Display(&todos)
	},
}

func init() {
	rootCmd.AddCommand(listaAllCMd)
}