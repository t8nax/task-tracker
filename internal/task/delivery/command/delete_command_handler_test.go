package command

import (
	"errors"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/t8nax/task-tracker/internal/task/usecase"
)

func TestDeleteExecute_ReturnsError_WhenArgumentsIsInvalid(t *testing.T) {
	ctl := gomock.NewController(t)
	uCase := usecase.NewMockTaskUseCase(ctl)

	handler := &DeleteCommandHandler{
		uCase: uCase,
	}

	messages, err := handler.Execute([]string{})

	assert.Nil(t, messages)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrInvalidArguments.Error())
}

func TestDeketeExecute_ReturnsError_WhenIDIsInvalid(t *testing.T) {
	ctl := gomock.NewController(t)
	uCase := usecase.NewMockTaskUseCase(ctl)

	handler := &DeleteCommandHandler{
		uCase: uCase,
	}

	ID := "123abc"

	messages, err := handler.Execute([]string{"/path/to/file", "delete", ID})

	assert.Nil(t, messages)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrUnableToParseTaskId.Error())
}

func TestDeleteExecute_ReturnsError_WhenUCaseReturnsError(t *testing.T) {
	ctl := gomock.NewController(t)
	uCase := usecase.NewMockTaskUseCase(ctl)

	ID := uint64(1)

	uCaseErr := errors.New("unable to delete tasks")

	uCase.EXPECT().DeleteTask(ID).Return(uCaseErr)

	handler := &DeleteCommandHandler{
		uCase: uCase,
	}

	messages, err := handler.Execute([]string{"/path/to/file", "delete", strconv.FormatUint(ID, 10)})

	assert.Nil(t, messages)
	assert.Error(t, err)
	assert.ErrorIs(t, err, uCaseErr)
}

func TestDeleteExecute_ReturnsEmptySlice_WhenExecutionIsSuccessful(t *testing.T) {
	ctl := gomock.NewController(t)
	uCase := usecase.NewMockTaskUseCase(ctl)

	ID := uint64(1)

	uCase.EXPECT().DeleteTask(ID).Return(nil)

	handler := &DeleteCommandHandler{
		uCase: uCase,
	}

	messages, err := handler.Execute([]string{"/path/to/file", "delete", strconv.FormatUint(ID, 10)})

	assert.Nil(t, err)
	assert.NotNil(t, messages)
	assert.Empty(t, messages)
}
