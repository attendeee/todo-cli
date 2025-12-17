package cmd

import (
	"fmt"
	"strings"
	"task/repository"

	"github.com/spf13/cobra"
)

func AddCathegory(cathegory string) {
	err := repository.AddCathegory(cathegory)
	if err != nil {
		panic(fmt.Errorf("Unable to add cathegory: %s\n", cathegory))
	}
	fmt.Println("Cathegory has been successfully added")

}

var cathegoryAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add cathegory",
	Run: func(cmd *cobra.Command, args []string) {

		cathegory := strings.Join(args, "")
		repository.AddCathegory(cathegory)

	},
}

func init() {
	cathegoryCmd.AddCommand(cathegoryAddCmd)

}
