package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/phillipahereza/tasks/db"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Long:  "Mark a task on your TODO list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.FetchTasks()

		var tasksText []string

		for _, task := range tasks {
			tasksText = append(tasksText, fmt.Sprintf("%d: %s", task.ID, task.Value))
		}

		prompt := promptui.Select{
			Label: "Select a task to complete",
			Items: tasksText,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		id := strings.Split(result, ":")[0]

		// convert ID to int

		idInt, err := strconv.Atoi(id)

		if err != nil {
			fmt.Printf("Failed to complete selected task %v\n", err)
			return
		}

		err = db.DeleteTask(idInt)

		fmt.Printf("You completed %q\n", result)

	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
