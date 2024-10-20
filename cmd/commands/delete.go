package commands

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"terminal-todos/internal/todo"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <index>",
	Short: "delete - delete a todo by id",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("todo id required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		todos, err := todo.Delete(id)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		} else {
			todo.Display(&todos)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
