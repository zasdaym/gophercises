package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zasdaym/gophercises/task/db"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("failed to get tasks.")
			return
		}

		fmt.Printf("ID\tTask\n")
		fmt.Printf("----------------------------------\n")
		for _, task := range tasks {
			fmt.Printf("%d\t%s\n", task.Key, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
