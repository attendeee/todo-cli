package cmd

import (
	"fmt"
	"task/repository"

	"github.com/spf13/cobra"
)

func init() {

	updateCathegoryCmd.Flags().Int("id", 0, "Set id of cathegory you want to update")
	updateCathegoryCmd.Flags().String("name", "", "Set new name of cathegory")

	cathegoryCmd.AddCommand(updateCathegoryCmd)
}

var updateCathegoryCmd = &cobra.Command{
	Use:   "update",
	Short: "Update cathegory",
	Run: func(cmd *cobra.Command, args []string) {

		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")

		if name == "" || id == 0 {
			fmt.Println("Pass id and name through flags")
			return
		}

		if id == 1 {
			fmt.Println("Unable to update default cathegory")
			return
		}

		repository.UpdateCathegory(id, name)
	},
}
