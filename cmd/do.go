package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use: "do",
	Short: "Mark a task on your TODO list as complete",
	Long: "Mark a task on your TODO list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		task := "review talk proposal"
		fmt.Printf("You have completed the \"%s\" task", task)
	},
}

func init(){
	rootCmd.AddCommand(doCmd)
}