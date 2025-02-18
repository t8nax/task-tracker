package command

import (
	"strconv"

	"github.com/t8nax/task-tracker/internal/task"
	"github.com/t8nax/task-tracker/internal/task/entity"
)

type MarkCommandHanlder struct {
	uCase  task.TaskUseCase
	status entity.Status
}

func (h *MarkCommandHanlder) Execute(args []string) ([]string, error) {
	if len(args) < 3 {
		return nil, ErrInvalidArguments
	}

	ID, err := strconv.ParseUint(args[2], 10, 64)

	if err != nil {
		return nil, ErrUnableToParseTaskId
	}

	_, err = h.uCase.UpdateTask(ID, h.status, "")

	return nil, err
}
