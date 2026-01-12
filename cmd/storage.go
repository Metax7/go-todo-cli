package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

func tasksFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("cannot determine user home directory")
	}

	dir := filepath.Join(home, ".todo")
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		panic("cannot create .todo directory in home")
	}

	return filepath.Join(dir, "tasks.json")
}

func loadTasks() ([]Task, error) {
	fileName := tasksFilePath()
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func saveTasks(tasks []Task) error {
	fileName := tasksFilePath()
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}
