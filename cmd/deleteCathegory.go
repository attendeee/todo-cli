package cmd

import (
	"task/repository"

	"github.com/spf13/cobra"
)

var deleteCathegoryCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete cathegory (Be careful. All tasks assigned to this cathegory will be deleted)",
	Run: func(cmd *cobra.Command, args []string) {

		id, _ := cmd.Flags().GetUint("id")

		repository.DeleteCathegory(int(id))
	},
}

func init() {
	deleteCathegoryCmd.Flags().Uint("id", 0, "Id of cathegory")
	cathegoryCmd.AddCommand(deleteCathegoryCmd)
}
