package cmd

import (
	"fmt"
	"github.com/phillipahereza/tasks/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Long:  "List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := db.FetchTasks()
		if err != nil {
			fmt.Println("Unable to fetch task list")
		}
		for i, t := range tasks {
			fmt.Printf("%d: %s\n", i+1, t.Value)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
