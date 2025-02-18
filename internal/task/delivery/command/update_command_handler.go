package command

import (
	"strconv"

	task "github.com/t8nax/task-tracker/internal/task"
)

type UpdateCommandHanlder struct {
	uCase task.TaskUseCase
}

func (h *UpdateCommandHanlder) Execute(args []string) ([]string, error) {
	if len(args) < 4 {
		return nil, ErrInvalidArguments
	}

	ID, err := strconv.ParseUint(args[2], 10, 64)

	if err != nil {
		return nil, ErrUnableToParseTaskId
	}

	description := args[3]

	_, err = h.uCase.UpdateTask(ID, "", description)

	return nil, err
}
