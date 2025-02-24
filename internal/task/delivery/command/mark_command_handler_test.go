package command

import (
	"errors"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/t8nax/task-tracker/internal/task/entity"
	"github.com/t8nax/task-tracker/internal/task/usecase"
)

func TestMarkExecute_ReturnsError_WhenArgumentsIsInvalid(t *testing.T) {
	ctl := gomock.NewController(t)
	uCase := usecase.NewMockTaskUseCase(ctl)

	handler := &MarkCommandHanlder{
		uCase: uCase,
	}

	messages, err := handler.Execute([]string{})

	assert.Nil(t, messages)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrInvalidArguments.Error())
}

func TestMarkExecute_ReturnsError_WhenIDIsInvalid(t *testing.T) {
	ctl := gomock.NewController(t)
	uCase := usecase.NewMockTaskUseCase(ctl)

	handler := &MarkCommandHanlder{
		uCase: uCase,
	}

	ID := "123abc"

	messages, err := handler.Execute([]string{"/path/to/file", "mark-done", ID})

	assert.Nil(t, messages)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrUnableToParseTaskId.Error())
}

func TestMarkExecute_ReturnsError_WhenUCaseReturnsError(t *testing.T) {
	ctl := gomock.NewController(t)
	uCase := usecase.NewMockTaskUseCase(ctl)

	ID := uint64(1)
	status := entity.StatusDone

	uCaseErr := errors.New("unable to update tasks")

	uCase.EXPECT().UpdateTask(ID, status, "").Return(nil, uCaseErr)

	handler := &MarkCommandHanlder{
		uCase:  uCase,
		status: status,
	}

	messages, err := handler.Execute([]string{"/path/to/file", "mark-done", strconv.FormatUint(ID, 10)})

	assert.Nil(t, messages)
	assert.Error(t, err)
	assert.ErrorIs(t, err, uCaseErr)
}

func TestMarkExecute_ReturnsNils_WhenExecutionIsSuccessful(t *testing.T) {
	ctl := gomock.NewController(t)
	uCase := usecase.NewMockTaskUseCase(ctl)

	status := entity.StatusDone
	ID := uint64(1)

	uCase.EXPECT().UpdateTask(ID, status, "").Return(&entity.Task{ID: ID}, nil)

	handler := &MarkCommandHanlder{
		uCase: uCase,
		status: status,
	}

	messages, err := handler.Execute([]string{"/path/to/file", "mark-done", strconv.FormatUint(ID, 10)})

	assert.Nil(t, err)
	assert.Nil(t, messages)
}