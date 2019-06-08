package cmd

import (
	"github.com/phillipahereza/tasks/db"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Long:  "Mark a task on your TODO list as complete",
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
			err := db.DoTask(id)
			if err != nil {
				fmt.Printf("Failed to mark task %d as complete\n", id)
			}
			fmt.Printf("You have completed the \"%d\" task\n", id)
		}

	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
