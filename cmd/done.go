package cmd

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark a task as done",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid task ID")
		}

		tasks, err := loadTasks()
		if err != nil {
			return err
		}

		found := false
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Done = true
				found = true
				break
			}
		}

		if !found {
			color.Red("❌ Task %d not found", id)
			return nil
		}

		if err := saveTasks(tasks); err != nil {
			return err
		}

		green := color.New(color.FgGreen, color.Bold).SprintFunc()
		fmt.Printf("%s Task %d marked as done!\n", green("🎉"), id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
