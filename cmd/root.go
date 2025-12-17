package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command = &cobra.Command{
	Use:   "task",
	Short: "A simple CLI task tool",
	Run: func(cmd *cobra.Command, args []string) {
		readCmd.Run(cmd, args)
	},
}

func RootCmd() *cobra.Command {
	return rootCmd
}
