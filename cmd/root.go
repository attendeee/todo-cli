package cmd

import (
	"task/repository"
	"task/table"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command = &cobra.Command{
	Use:   "task",
	Short: "A simple CLI task tool",
	Run: func(cmd *cobra.Command, args []string) {
		readCmd.Run(cmd, args)
	},
}

var cathegoryCmd *cobra.Command = &cobra.Command{
	Use:   "cathegory",
	Short: "Subset of commands to work with cathegories",
	Run: func(cmd *cobra.Command, args []string) {
		readCathegoryCmd.Run(cmd, args)
	},
}

func init() {
	repository.InitDb()
	table.Init()

	rootCmd.AddCommand(cathegoryCmd)
}

func RootCmd() *cobra.Command {

	return rootCmd
}
