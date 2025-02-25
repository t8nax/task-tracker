package command

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/t8nax/task-tracker/internal/task/entity"
	"github.com/t8nax/task-tracker/internal/task/usecase"
)

func TestListExecute_ReturnsError_WhenArgumentsIsInvalid(t *testing.T) {
	ctl := gomock.NewController(t)
	uCase := usecase.NewMockTaskUseCase(ctl)

	handler := &ListCommandHanlder{
		uCase: uCase,
	}

	messages, err := handler.Execute([]string{})

	assert.Nil(t, messages)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrInvalidArguments.Error())
}

func TestListExecute_ReturnsError_WhenStatusIsInvalid(t *testing.T) {
	ctl := gomock.NewController(t)
	uCase := usecase.NewMockTaskUseCase(ctl)

	handler := &ListCommandHanlder{
		uCase: uCase,
	}

	status := "invalid_status"
	messages, err := handler.Execute([]string{"path/to/file", "list", status})

	assert.Nil(t, messages)
	assert.Error(t, err)
	assert.EqualError(t, err, entity.GetErrInvalidStatus(status).Error())
}

func TestListExecute_ReturnsError_WhenUCaseReturnsError(t *testing.T) {
	ctl := gomock.NewController(t)
	uCase := usecase.NewMockTaskUseCase(ctl)

	uCaseErr := errors.New("unable to get tasks")

	uCase.EXPECT().GetAllTasks(entity.StatusNone).Return(nil, uCaseErr)

	handler := &ListCommandHanlder{
		uCase: uCase,
	}

	messages, err := handler.Execute([]string{"path/to/file", "list"})

	assert.Nil(t, messages)
	assert.Error(t, err)
	assert.EqualError(t, err, uCaseErr.Error())
}

func TestListExecute_ReturnsEmptySlice_WhenNoTasksInList(t *testing.T) {
	ctl := gomock.NewController(t)
	uCase := usecase.NewMockTaskUseCase(ctl)

	uCase.EXPECT().GetAllTasks(entity.StatusNone).Return([]entity.Task{}, nil)

	handler := &ListCommandHanlder{
		uCase: uCase,
	}

	messages, err := handler.Execute([]string{"path/to/file", "list"})

	assert.Nil(t, err)
	assert.NotNil(t, messages)
	assert.Empty(t, messages)
}

func TestListExecute_ReturnsTaskDescriptions_WhenThereIsTasksInList(t *testing.T) {
	ctl := gomock.NewController(t)
	uCase := usecase.NewMockTaskUseCase(ctl)

	tasks := []entity.Task{
		{
			ID: 1,
		},
		{
			ID: 2,
		},
	}

	uCase.EXPECT().GetAllTasks(entity.StatusNone).Return(tasks, nil)

	handler := &ListCommandHanlder{
		uCase: uCase,
	}

	messages, err := handler.Execute([]string{"path/to/file", "list"})

	assert.Nil(t, err)
	assert.NotNil(t, messages)

	for i, msg := range messages {
		assert.Contains(t, msg, fmt.Sprintf("ID: %d", tasks[i].ID))
	}
}
