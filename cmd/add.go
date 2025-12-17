package cmd

import (
	"fmt"
	"task/repository"

	"github.com/spf13/cobra"
)

func addTask(catheogry int16, task string, notes string) int64 {
	id, err := repository.AddTask(catheogry, task, notes)
	if err != nil {
		panic(fmt.Errorf("Unable to add task: %s\n", err))
	}

	return id

}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task",
	Run: func(cmd *cobra.Command, args []string) {
		c, _ := cmd.Flags().GetInt16("cathegory")
		t, _ := cmd.Flags().GetString("topic")
		n, _ := cmd.Flags().GetString("notes")

		addTask(c, t, n)

		fmt.Println("Task has been successfully added")

	},
}

func init() {
	addCmd.Flags().Int16P("cathegory", "c", 1, "Set cathegory of task")
	addCmd.Flags().StringP("topic", "t", "", "Set topic of task")
	addCmd.Flags().StringP("notes", "n", "", "Set additional notes for task")

	RootCmd().AddCommand(addCmd)
}
