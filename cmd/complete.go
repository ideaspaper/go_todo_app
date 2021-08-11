package cmd

import "github.com/spf13/cobra"

var completeId *int

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "mark task as complete by id",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		todoController.Complete(completeId)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
	completeId = completeCmd.Flags().IntP("id", "i", 0, "task id")
	completeCmd.MarkFlagRequired("id")
}
