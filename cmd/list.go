package cmd

import (
	"fmt"
	"project-app-todo-list-cli/service"
	"project-app-todo-list-cli/utils"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "show all list task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := service.DisplayTodoList()

		headers := []string{"No", "Task", "Status", "Priority"}
		rows := [][]string{}

		for i, t := range tasks {
			rows = append(rows, []string{
				fmt.Sprintf("%d", i+1),
				t.Title,
				t.Status,
				t.Priority,
			})
		}

		utils.PrintTable(headers, rows)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
