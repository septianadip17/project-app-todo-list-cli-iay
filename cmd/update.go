package cmd

import (
	"fmt"
	"project-app-todo-list-cli/service"
	"project-app-todo-list-cli/utils"

	"github.com/spf13/cobra"
)

var updateIndex int
var updateStatus string

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update task status",
	Run: func(cmd *cobra.Command, args []string) {
		if updateIndex <= 0 {
			utils.PrintError("index harus lebih dari 0")
			return
		}

		if updateStatus == "" {
			utils.PrintError("status harus diisi")
			return
		}

		if !utils.IsValidStatus(updateStatus) {
			utils.PrintError("status hanya boleh new, progress, completed")
			return
		}

		err := service.UpdateStatus(updateIndex-1, updateStatus)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Task status updated successfully")
	},
}

func init() {
	updateCmd.Flags().IntVar(&updateIndex, "index", 0, "nomor task yang ingin diupdate")
	updateCmd.Flags().StringVar(&updateStatus, "status", "", "status task")
	rootCmd.AddCommand(updateCmd)
}
