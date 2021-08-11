package cmd

import "github.com/spf13/cobra"

var findId *int

var findIdCmd = &cobra.Command{
	Use:   "findId",
	Short: "find task by id",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		todoController.FindById(findId)
	},
}

func init() {
	rootCmd.AddCommand(findIdCmd)
	findId = findIdCmd.Flags().IntP("id", "i", 0, "task id")
	findIdCmd.MarkFlagRequired("id")
}
