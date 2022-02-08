package cmd

import (
	"fmt"
	"os"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("can't add empty task to task list")
		} else {
			task := strings.Join(args, " ")
			id, err := db.CreateTask(task)
			if err != nil {
				fmt.Println("Something went wrong !", err.Error())
				os.Exit(1)
			}
			fmt.Printf("Added \"%s \" to your task list in row %d.", task, id)
		}
	},
}
