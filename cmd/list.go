package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Long:  "List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := []string{"Cook food", "Eat food", "Wash plates"}
		for i, t := range tasks {
			fmt.Printf("%d: %s\n", i+1, t)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
