package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/phillipahereza/tasks/db"
	"github.com/spf13/cobra"
	"strings"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Long:  "Mark a task on your TODO list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.FetchTasks()

		if len(tasks) < 1 {
			fmt.Println("You have no incomplete tasks")
			return
		}

		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "\U00002714 {{ .Value | cyan }}",
			Inactive: "  {{ .Value | cyan }}",
			Selected: "\U00002714 {{ .Value | red | cyan }}",
			Details: `
--------- Tasks ----------
{{ "Task:" | faint }}	{{ .Value }}
{{ "Started:" | faint }}	{{ .CreatedAt }}`,
		}

		searcher := func(input string, index int) bool {
			task := tasks[index]
			name := strings.Replace(strings.ToLower(task.Value), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)

			return strings.Contains(name, input)
		}


		prompt := promptui.Select{
			Label: "Select a task to complete",
			Items: tasks,
			Templates: templates,
			Searcher:  searcher,
			Size:      9,
		}

		i, _, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}


		err = db.DoTask(tasks[i].ID)

	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
