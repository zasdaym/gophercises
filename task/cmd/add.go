package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zasdaym/gophercises/task/db"
)

// addCmd represents the add command.
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, "")
		id, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Failed to create task.")
			return
		}
		fmt.Printf("Added %s to your task list with id %d.\n", task, id)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
