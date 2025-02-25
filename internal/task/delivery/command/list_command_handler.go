package command

import (
	"fmt"

	task "github.com/t8nax/task-tracker/internal/task"
	"github.com/t8nax/task-tracker/internal/task/entity"
)

type ListCommandHanlder struct {
	uCase task.TaskUseCase
}

func (h *ListCommandHanlder) Execute(args []string) ([]string, error) {
	if len(args) < 2 {
		return nil, ErrInvalidArguments
	}

	status := entity.StatusNone
	var err error

	if len(args) > 2 {
		status, err = entity.ParseStatus(args[2])

		if err != nil {
			return nil, err
		}
	}

	tasks, err := h.uCase.GetAllTasks(status)

	if err != nil {
		return nil, err
	}

	const template = "01-02-2006 15:04:05"
	res := make([]string, 0, len(tasks))

	for _, task := range tasks {
		res = append(res, fmt.Sprintf("ID: %d; Description: %s; Status: %s; Created at: %s; Updated at: %s",
			task.ID, task.Description, task.Status, task.CreatedAt.Format(template), task.UpdatedAt.Format(template)))
	}

	return res, nil
}
