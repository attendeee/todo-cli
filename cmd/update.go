package cmd

import (
	"fmt"
	"strconv"
	"task/repository"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd().AddCommand(completeCmd)
	RootCmd().AddCommand(incompleteCmd)

	RootCmd().AddCommand(updateCmd)

	updateCmd.Flags().Int("id", 0, "Update task data")
	updateCmd.Flags().StringP("topic", "t", "", "Update topic data")
	updateCmd.Flags().StringP("notes", "n", "", "Update notes data")

}

func updateTaskRecord(task repository.Task) {
	repository.UpdateTask(task.Id, task.Topic, task.Notes)
}

func incomplete(args []string) {
	for _, i := range args {
		id, err := strconv.Atoi(i)
		if err != nil {
			panic("This is not number")
		}

		repository.IncompleteTask(id)

		fmt.Printf("Tasks %d has been marked as incompleted\n", id)
	}
}
func complete(args []string) {
	for _, i := range args {
		id, err := strconv.Atoi(i)
		if err != nil {
			panic("This is not number")
		}

		repository.CompleteTask(id)

		fmt.Printf("Tasks %d has been marked as completed\n", id)
	}
}

func update(cmd *cobra.Command) {

	id, _ := cmd.Flags().GetInt("id")
	topic, _ := cmd.Flags().GetString("topic")
	notes, _ := cmd.Flags().GetString("notes")

	t := repository.Task{Id: id, Topic: topic, Notes: notes}
	updateTaskRecord(t)

	fmt.Printf("Task with id %d is updated\n", id)
}

var incompleteCmd = &cobra.Command{
	Use:   "undo",
	Short: "Complete task",
	Run: func(cmd *cobra.Command, args []string) {
		incomplete(args)

	},
}
var completeCmd = &cobra.Command{
	Use:   "do",
	Short: "Complete task",
	Run: func(cmd *cobra.Command, args []string) {
		complete(args)

	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update task",
	Run: func(cmd *cobra.Command, args []string) {
		update(cmd)

	},
}
