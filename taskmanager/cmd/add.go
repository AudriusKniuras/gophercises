package cmd

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add task",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("task: \"%s\" added to the list\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
