package cmd

import (
	"fmt"
	"github.com/phillipahereza/tasks/db"
	"github.com/spf13/cobra"
)

var status string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Long:  "List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := db.FetchTasks()
		if err != nil {
			fmt.Println("Unable to fetch task list")
		}
		for _, t := range tasks {
			fmt.Printf("%d: %s\n", t.ID, t.Value)
		}
		
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
