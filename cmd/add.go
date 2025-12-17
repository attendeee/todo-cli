package cmd

import (
	"fmt"
	"task/repository"

	"github.com/spf13/cobra"
)

func addTask(task string, notes string) int64 {
	id, err := repository.AddTask(task, notes)
	if err != nil {
		panic(fmt.Errorf("Unable to add task: %s\n", err))
	}

	fmt.Println("Task has been successfully added")
	return id

}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task",
	Run: func(cmd *cobra.Command, args []string) {
		t, _ := cmd.Flags().GetString("topic")
		n, _ := cmd.Flags().GetString("notes")

		addTask(t, n)

	},
}

func init() {
	addCmd.Flags().StringP("topic", "t", "", "Set topic of task")
	addCmd.Flags().StringP("notes", "n", "", "Set additional notes for task")

	RootCmd().AddCommand(addCmd)
}
