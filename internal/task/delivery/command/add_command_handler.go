package command

import (
	"fmt"

	"github.com/t8nax/task-tracker/internal/task"
)

type AddCommandHanlder struct {
	uCase task.TaskUseCase
}

func (h *AddCommandHanlder) Execute(args []string) ([]string, error) {
	if len(args) < 3 {
		return nil, ErrInvalidArguments
	}

	description := args[2]
	task, err := h.uCase.AddTask(description)

	if err != nil {
		return nil, err
	}

	return []string{fmt.Sprintf("Task added successfully (ID: %d)", task.ID)}, nil
}
