package cmd

import (
	"fmt"
	"task/repository"
	"task/table"

	"github.com/spf13/cobra"
)

func CathegoryReadAndPrint() {
	cathegories, _ := repository.ReadCathegory()

	if len(*cathegories) == 0 {
		fmt.Println("There is no cathegories")
		return
	}

	fmt.Println("Cathegories has been successfully read")

	table.PrintCathegoryTable(*cathegories)
}

var readCathegoryCmd = &cobra.Command{
	Use:   "read",
	Short: "Read cathegories",
	Run: func(cmd *cobra.Command, args []string) {

		CathegoryReadAndPrint()

	},
}

func init() {
	cathegoryCmd.AddCommand(readCathegoryCmd)
}
