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
		cathegoryName, _ := cmd.Flags().GetString("cathegory")
		if cathegoryName == "" {
			cathegoryName = "default"
		}

		tasks, err := repository.ReadAssignedTasks(cathegoryName)
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
	readCmd.Flags().StringP("cathegory", "c", "default", "Read tasks assigned to cathegory name")

	RootCmd().AddCommand(readCmd)
}
