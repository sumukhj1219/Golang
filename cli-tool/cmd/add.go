package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/myapp/models"
)

var addTask = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the list",
	Long:  "This command allows you to add a new task to your task list.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := models.LoadTasks()
		if err != nil {
			fmt.Println("Error in loading tasks ❌", err)
			return
		}
		newTask := models.Todo{
			Id:          string(len(tasks) + 1),
			Title:       args[0],
			Created:     time.Now(),
			Description: args[1],
			Status:      false,
		}
		tasks = append(tasks, newTask)
		err = models.SaveTasks(tasks)
		if err != nil {
			fmt.Println("Error in saving tasks ❌")
		}
		fmt.Println("Task added ✅", newTask)
	},
}

func init() {
	rootCmd.AddCommand(addTask)
}
