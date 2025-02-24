package command

import (
	"fmt"

	"github.com/t8nax/task-tracker/internal/task"
)

func getTaskAddedMessage(ID uint64) string {
	return fmt.Sprintf("Task added successfully (ID: %d)", ID)
}

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

	return []string{getTaskAddedMessage(task.ID)}, nil
}
