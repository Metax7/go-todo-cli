package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := loadTasks()
		if err != nil {
			return err
		}

		if len(tasks) == 0 {
			color.Yellow("⚠️  No tasks found.")
			return nil
		}

		header := color.New(color.Bold, color.FgHiCyan).SprintFunc()
		fmt.Println(header("🗒  Your To-Do List:"))
		fmt.Println()

		for _, t := range tasks {
			if t.Done {
				gray := color.New(color.Faint, color.FgHiBlack).SprintFunc()
				check := color.New(color.FgGreen).Sprint("✔")
				fmt.Printf("%s %2d. %s\n", check, t.ID, gray(t.Text))
			} else {
				bullet := color.New(color.FgYellow).Sprint("•")
				fmt.Printf("%s %2d. %s\n", bullet, t.ID, t.Text)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
