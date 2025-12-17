package cmd

import (
	"task/repository"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete task",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetUint("id")
		repository.DeleteTask(int(id))
	},
}

func init() {
	deleteCmd.Flags().Uint("id", 0, "Id of task")
	RootCmd().AddCommand(deleteCmd)
}
