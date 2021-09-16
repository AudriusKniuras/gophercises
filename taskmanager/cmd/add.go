package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/AudriusKniuras/gophercises/taskmanager/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add task",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		id, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Could not add task: ", err)
			os.Exit(1)
		}
		fmt.Printf("task %d: \"%s\" added to the list\n", id, task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
