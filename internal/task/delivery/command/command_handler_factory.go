package command

import (
	"fmt"

	"github.com/t8nax/task-tracker/internal/task"
	"github.com/t8nax/task-tracker/internal/task/entity"
)

type CommandHandlerFactory struct {
	uCase task.TaskUseCase
}

func NewCommandHandlerFactory(uCase task.TaskUseCase) *CommandHandlerFactory {
	if uCase == nil {
		panic("service must not be nil")
	}

	return &CommandHandlerFactory{
		uCase,
	}
}

func (f *CommandHandlerFactory) GetHandler(cmd Command) (CommandHandler, error) {
	switch cmd {
	case СommandList:
		return &ListCommandHanlder{f.uCase}, nil
	case СommandAdd:
		return &AddCommandHanlder{f.uCase}, nil
	case СommandMarkDone:
		return &MarkCommandHanlder{f.uCase, entity.StatusDone}, nil
	case СommandMarkInProgress:
		return &MarkCommandHanlder{f.uCase, entity.StatusInProgress}, nil
	case СommandUpdate:
		return &UpdateCommandHanlder{f.uCase}, nil
	default:
		return nil, fmt.Errorf("unknown command: %s", cmd)
	}
}
