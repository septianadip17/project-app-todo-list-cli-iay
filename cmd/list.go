package cmd

import (
	"fmt"
	"project-app-todo-list-cli/service"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "show all list task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := service.DisplayTodoList()

		fmt.Println("No | Task | Status | Priority")
		for i, t := range tasks {
			fmt.Printf("%d | %s | %s | %s\n", i+1, t.Title, t.Status, t.Priority)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
