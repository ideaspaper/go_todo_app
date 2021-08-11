package cmd

import "github.com/spf13/cobra"

var sortFlag *string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("sort") {
			todoController.List(sortFlag)
		} else {
			todoController.List(nil)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	sortFlag = listCmd.Flags().StringP("sort", "s", "", "list all tasks alphabetically in [asc|desc] order")
}
