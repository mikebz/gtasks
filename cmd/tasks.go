/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/mikebz/gtasks/internal/tasks"

	"github.com/spf13/cobra"
)

var (
	hiddenFlag    bool
	completedFlag bool
	assignedFlag  bool
)

// tasksCmd represents the tasks command
var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Show the list of tasks",
	Long: `Show a list of tasks based on the API:
https://developers.google.com/tasks/reference/rest/v1/tasks/list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := tasks.Tasks(hiddenFlag,
			completedFlag, assignedFlag)
		if err != nil {
			return err
		}
		for _, task := range tasks {
			println(task)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(tasksCmd)
	tasksCmd.PersistentFlags().BoolVarP(&hiddenFlag, "hidden", "h", false, "Show hidden tasks")
	tasksCmd.PersistentFlags().BoolVarP(&completedFlag, "completed", "c", false, "Show completed tasks")
	tasksCmd.PersistentFlags().BoolVarP(&assignedFlag, "assigned", "a", true, "Show assigned tasks")
	tasksCmd.Flags()
}
