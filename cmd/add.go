package cmd

import "github.com/spf13/cobra"

var taskName *string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add new task",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		todoController.Add(taskName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	taskName = addCmd.Flags().StringP("task", "t", "", "your task")
	addCmd.MarkFlagRequired("task")
}
