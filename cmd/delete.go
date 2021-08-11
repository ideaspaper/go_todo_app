package cmd

import "github.com/spf13/cobra"

var deleteId *int

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete task by id",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		todoController.Delete(deleteId)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteId = deleteCmd.Flags().IntP("id", "i", 0, "task id")
	deleteCmd.MarkFlagRequired("id")
}
