package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		text := args[0]
		tasks, err := loadTasks()
		if err != nil {
			return err
		}

		id := len(tasks) + 1
		tasks = append(tasks, Task{ID: id, Text: text})
		if err := saveTasks(tasks); err != nil {
			return err
		}

		green := color.New(color.FgGreen).SprintFunc()
		fmt.Printf("%s Task added: %s\n", green("✅"), text)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
