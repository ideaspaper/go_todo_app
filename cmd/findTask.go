package cmd

import "github.com/spf13/cobra"

var findTask *string

var findTaskCmd = &cobra.Command{
	Use:     "find",
	Short:   "find task by string",
	Aliases: []string{"findTask"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		todoController.FindByTask(findTask)
	},
}

func init() {
	rootCmd.AddCommand(findTaskCmd)
	findTask = findTaskCmd.Flags().StringP("task", "t", "", "task id")
	findTaskCmd.MarkFlagRequired("task")
}
