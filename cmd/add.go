package cmd

import (
	"fmt"
	"strings"

	"github.com/phillipahereza/tasks/db"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Long:  "Add a new task to your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		content := strings.Join(args, " ")
		savedTask, err := db.CreateTask(content)
		if err != nil {
			fmt.Printf("Error adding task: %v\n", err)
		} else {
			fmt.Printf("Added \"%s\" to your task list\n", savedTask.Value)
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
