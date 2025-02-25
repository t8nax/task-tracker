package command

import (
	"strconv"

	"github.com/t8nax/task-tracker/internal/task"
)

type DeleteCommandHandler struct {
	uCase task.TaskUseCase
}

func (h *DeleteCommandHandler) Execute(args []string) ([]string, error) {
	if len(args) < 3 {
		return nil, ErrInvalidArguments
	}

	ID, err := strconv.ParseUint(args[2], 10, 64)

	if err != nil {
		return nil, ErrUnableToParseTaskId
	}

	err = h.uCase.DeleteTask(ID)

	if err != nil {
		return nil, err
	}

	return []string{}, nil
}
