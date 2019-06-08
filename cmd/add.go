package cmd

import (
	"strings"
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Long: "Add a new task to your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" to your task list\n", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
