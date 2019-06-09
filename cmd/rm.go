package cmd

import (
	"fmt"
	"github.com/phillipahereza/tasks/db"
	"github.com/spf13/cobra"
	"strconv"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete a task",
	Long:  "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("failed to parse argument \"%s\"\n", arg)
			} else {
				ids = append(ids, id)
			}
		}
		for _, id := range ids {
			err := db.DeleteTask(id)
			if err != nil {
				fmt.Printf("Failed to delete task %d as complete\n", id)
			}
			fmt.Printf("You have deleted the \"%d\" task\n", id)
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
