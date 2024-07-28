/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/mikebz/gtasks/internal/tasks"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listsCmd = &cobra.Command{
	Use:   "lists",
	Short: "Show the taskslists",
	Long: `Show the tasklists:  For more documentation about the
output please see: https://developers.google.com/tasks/reference/rest/v1/tasklists/list
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tasklists, err := tasks.Lists()
		if err != nil {
			return err
		}
		for _, it := range tasklists {
			println(it)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listsCmd)
}
