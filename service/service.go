package service

import (
	"errors"
	"project-app-todo-list-cli/data"
	"project-app-todo-list-cli/model"
)

func init() {
	_ = data.Load()
}

func DisplayTodoList() []model.Task {
	return data.Tasks
}

func UpdateStatus(idx int, status string) error {
	if idx < 0 || idx >= len(data.Tasks) {
		return errors.New("index di luar jangkauan")
	}

	data.Tasks[idx].Status = status
	return data.Save()
}
