package command

import (
	"fmt"

	task "github.com/t8nax/task-tracker/internal/task"
)

type ListCommandHanlder struct {
	uCase task.TaskUseCase
}

func (h *ListCommandHanlder) Execute(args []string) ([]string, error) {
	tasks, err := h.uCase.GetAllTasks()

	if err != nil {
		return nil, err
	}

	const template = "01-02-2026 15:04:05"
	res := make([]string, len(tasks))

	for _, task := range tasks {
		res = append(res, fmt.Sprintf("ID: %d; Description: %s; Status: %s; Created at: %s; Updated at: %s",
			task.ID, task.Description, task.Status, task.CreatedAt.Format(template), task.UpdatedAt.Format(template)))
	}

	return res, nil
}
