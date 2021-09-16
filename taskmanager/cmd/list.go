package cmd

import (
	"fmt"
	"os"

	"github.com/AudriusKniuras/gophercises/taskmanager/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Unable to get tasks: ", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("Task list empty")
			return
		}
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
