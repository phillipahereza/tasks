package cmd

import (
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
		for _, task := range ids {
			fmt.Printf("You have completed the \"%d\" task\n", task)
		}

	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
