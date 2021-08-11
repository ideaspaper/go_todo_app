package cmd

import "github.com/spf13/cobra"

var uncompleteId *int

var uncompleteCmd = &cobra.Command{
	Use:   "uncomplete",
	Short: "mark a task as complete by id",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		todoController.Uncomplete(uncompleteId)
	},
}

func init() {
	rootCmd.AddCommand(uncompleteCmd)
	uncompleteId = uncompleteCmd.Flags().IntP("id", "i", 0, "task id")
	uncompleteCmd.MarkFlagRequired("id")
}
