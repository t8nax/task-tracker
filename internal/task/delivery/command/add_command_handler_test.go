package command

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/t8nax/task-tracker/internal/task/entity"
	"github.com/t8nax/task-tracker/internal/task/usecase"
)

func TestAddExecute_ReturnsError_WhenArgumentsIsInvalid(t *testing.T) {
	ctl := gomock.NewController(t)
	uCase := usecase.NewMockTaskUseCase(ctl)

	handler := &AddCommandHanlder{
		uCase: uCase,
	}

	messages, err := handler.Execute([]string{})

	assert.Nil(t, messages)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrInvalidArguments.Error())
}

func TestAddExecute_ReturnsError_WhenUCaseReturnsError(t *testing.T) {
	ctl := gomock.NewController(t)
	uCase := usecase.NewMockTaskUseCase(ctl)

	uCaseErr := errors.New("unable to get tasks")
	description := "Description"

	uCase.EXPECT().AddTask(description).Return(nil, uCaseErr)

	handler := &AddCommandHanlder{
		uCase: uCase,
	}

	messages, err := handler.Execute([]string{"/path/to/file", "add", description})

	assert.Nil(t, messages)
	assert.Error(t, err)
	assert.EqualError(t, err, uCaseErr.Error())
}

func TestAddExecute_ReturnsMessage_WhenExecutionIsSuccessful(t *testing.T) {
	ctl := gomock.NewController(t)
	uCase := usecase.NewMockTaskUseCase(ctl)

	description := "Description"
	ID := uint64(1)

	uCase.EXPECT().AddTask(description).Return(&entity.Task{ID: ID}, nil)

	handler := &AddCommandHanlder{
		uCase: uCase,
	}

	messages, err := handler.Execute([]string{"/path/to/file", "add", description})

	assert.Nil(t, err)
	assert.NotNil(t, messages)
	assert.Equal(t, 1, len(messages))
	assert.Equal(t, getTaskAddedMessage(ID), messages[0])
}