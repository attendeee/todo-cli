package cmd

import (
	"fmt"
	"task/repository"
	"task/table"

	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read tasks",
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := repository.ReadTasks()

		if err != nil {
			panic(fmt.Errorf("Unable to read tasks: %s\n", err))
		}

		if len(*tasks) == 0 {
			fmt.Println("There is no tasks")
			return
		}

		table.PrintTaskTable(*tasks)
	},
}

func init() {
	RootCmd().AddCommand(readCmd)
}
