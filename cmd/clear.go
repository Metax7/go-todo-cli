package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Remove all completed tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := loadTasks()
		if err != nil {
			return err
		}

		var remaining []Task
		for _, t := range tasks {
			if !t.Done {
				remaining = append(remaining, t)
			}
		}

		if err := saveTasks(remaining); err != nil {
			return err
		}

		cyan := color.New(color.FgHiCyan).SprintFunc()
		fmt.Printf("%s Cleared completed tasks!\n", cyan("🧽"))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
