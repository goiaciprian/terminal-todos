package commands

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"terminal-todos/internal/todo"

	"github.com/spf13/cobra"
)

var title, desc string

var updateCmd = &cobra.Command{
	Use: "update <index>",
	Short: "update - update the title or the description of a todo",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("todo id missing")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "id is not a number %s\n", err)
			return
		}

		todos, err := todo.Update(id, title, desc)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			todo.Display(&todos)
		}

	},
}

func init() {
	updateCmd.Flags().StringVarP(&title, "title", "t", "", "new title value")
	updateCmd.Flags().StringVarP(&desc, "desc", "d", "", "new description value")
	rootCmd.AddCommand(updateCmd)
}
