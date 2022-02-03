package cmd

import (
	"fmt"
	"strings"

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
			fmt.Printf("Added \"%s \" to your task list.", task)
		}
	},
}
