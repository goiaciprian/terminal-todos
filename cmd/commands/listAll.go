package commands

import (
	"errors"
	"fmt"
	"os"
	"terminal-todos/internal/todo"

	"github.com/spf13/cobra"
)

var (
	all, completed, uncompleted bool
	sort                        string
)

var listaAllCMd = &cobra.Command{
	Use:   "list",
	Short: "list - display all todos",
	Args: func(cmd *cobra.Command, args []string) error {
		if (all && uncompleted && completed) ||
			(all && completed) ||
			(all && uncompleted) ||
			(completed && uncompleted) {
			return errors.New("only one of the -a -c -u is accepted")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, argv []string) {
		qtype := "uncompleted"
		if all {
			qtype = "all"
		} else if completed {
			qtype = "completed"
		}

		todos, err := todo.GetAll(qtype, sort)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting the todos %s", err)
		}

		todo.Display(&todos)
	},
}

func init() {
	listaAllCMd.Flags().BoolVarP(&all, "all", "a", false, "get all todos")
	listaAllCMd.Flags().BoolVarP(&completed, "completed", "c", false, "get all completed todos")
	listaAllCMd.Flags().BoolVarP(&uncompleted, "uncompleted", "u", false, "get all uncompleted todos")
	listaAllCMd.Flags().StringVarP(&sort, "sort", "s", "desc", "sort")
	rootCmd.AddCommand(listaAllCMd)
}
