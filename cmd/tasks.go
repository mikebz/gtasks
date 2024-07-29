/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	t "github.com/mikebz/gtasks/internal/tasks"

	"github.com/spf13/cobra"
)

var (
	allFlag       bool
	completedFlag bool
	assignedFlag  bool
	verboseFlag   bool
)

// tasksCmd represents the tasks command
var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Show the list of tasks",
	Long: `Show a list of tasks based on the API:
https://developers.google.com/tasks/reference/rest/v1/tasks/list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := t.Tasks(allFlag,
			completedFlag, assignedFlag)
		if err != nil {
			return err
		}
		for _, task := range tasks {
			if !verboseFlag {
				println(t.TaskToLine(task))
			} else {
				println(t.TaskVerbose(task))
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(tasksCmd)
	tasksCmd.Flags().BoolVar(&allFlag, "all", false, "Show all, including hidden")
	tasksCmd.Flags().BoolVarP(&completedFlag, "completed", "c", false, "Show completed tasks")
	tasksCmd.Flags().BoolVarP(&assignedFlag, "assigned", "a", true, "Show assigned tasks")
	tasksCmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "Show assigned tasks")
	tasksCmd.Flags()
}
