package cmd

import (
	"fmt"
	"strconv"
	"task/repository"

	"github.com/spf13/cobra"
)

func deleteTasks(args []string) {
	for _, id := range args {
		i, err := strconv.Atoi(id)
		if err != nil {
			fmt.Printf("This is seems to be not integer\n")
		}

		repository.DeleteTask(i)
		fmt.Printf("Task with id %d is deleted\n", i)
	}
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete task",
	Run: func(cmd *cobra.Command, args []string) {
		deleteTasks(args)
	},
}

func init() {
	RootCmd().AddCommand(deleteCmd)
}
