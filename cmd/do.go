package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete ✔",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Failed to parse the argument : %s \n", arg)
			} else {
				ids = append(ids, id)
			}
		}
		if len(ids) != 0 {
			for _, id := range ids {
				db.DeleteTask(id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)

}
