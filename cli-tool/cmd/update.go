/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/myapp/models"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Command used to update status of particular task",

	Run: func(cmd *cobra.Command, args []string) {
		_, err := models.UpdateTasks(args[0])
		if err != nil {
			fmt.Println("Error in updating task")
		}
		fmt.Println("Task updated successfully ✅")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

}
