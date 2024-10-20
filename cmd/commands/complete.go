package commands

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"terminal-todos/internal/todo"

	"github.com/spf13/cobra"
)

var uncomplete bool

var completeCmd = &cobra.Command{
	Use:   "complete <index>",
	Short: "complete - updates the todo as completed",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("todo id required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, argv []string) {
		id, err := strconv.ParseInt(argv[0], 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "id is not valid: %s\n", err)
			return
		}

		if uncomplete {
			todos, err := todo.Uncomplete(id)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			} else {
				todo.Display(&todos)	
			}
			return
		}
		todos, err := todo.Complete(id)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			todo.Display(&todos)
		}
	},
}

func init() {
	completeCmd.Flags().BoolVarP(&uncomplete, "uncomplete", "u", false, "set todo as uncomplete")
	rootCmd.AddCommand(completeCmd)
}
