package cmd

import (
	"fmt"
	"github.com/phillipahereza/tasks/db"
	"github.com/spf13/cobra"
	"time"
)

var hours int

var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "Shows a list of completed tasks",
	Long:  "Shows a list of completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		timeDelta := time.Duration(hours) * time.Hour
		tasks, err := db.FetchCompletedTasks(timeDelta)
		if err != nil {
			fmt.Printf("Unable to fetch task list: %v", err)
		}
		for _, t := range tasks {
			fmt.Printf("%d: %s\n", t.ID, t.Value)
		}
	},
}

func init() {
	completedCmd.Flags().IntVarP(&hours, "time", "t", 24, "tasks completed in the last X hours")
	rootCmd.AddCommand(completedCmd)

}
