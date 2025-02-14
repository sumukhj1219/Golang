/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/myapp/models"
)

var list = &cobra.Command{
	Use:   "list",
	Short: "List the available tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := models.LoadTasks()
		if err != nil {
			fmt.Println("Error in loading tasks")
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Title", "Description", "Created", "Status"})

		for i, task := range tasks {
			status := "🕛 Pending"
			if task.Status {
				status = "✔ Completed"
			}
			table.Append([](string){
				string(i),
				task.Title,
				task.Description,
				task.Created.Format("2006-01-02"),
				status,
			})
		}
		table.SetBorder(true)
		table.SetRowLine(true)
		table.SetCenterSeparator("|")
		table.SetColumnSeparator("|")
		table.SetRowSeparator("-")
		table.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
		table.SetAlignment(tablewriter.ALIGN_LEFT)

		table.Render()

	},
}

func init() {
	rootCmd.AddCommand(list)
}
