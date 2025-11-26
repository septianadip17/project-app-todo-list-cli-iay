package data

import (
	"encoding/json"
	"os"
	"project-app-todo-list-cli/model"
)

var Tasks []model.Task
var filePath = "data/tasks.json"

func Load() error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			Tasks = []model.Task{}
			return nil
		}
		return err
	}

	return json.Unmarshal(file, &Tasks)
}

func Save() error {
	b, err := json.MarshalIndent(Tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, b, 0644)
}
