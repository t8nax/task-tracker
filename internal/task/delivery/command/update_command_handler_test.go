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

func TestUpdateExecute_ReturnsError_WhenArgumentsIsInvalid(t *testing.T) {
	ctl := gomock.NewController(t)
	uCase := usecase.NewMockTaskUseCase(ctl)

	handler := &UpdateCommandHanlder{
		uCase: uCase,
	}

	messages, err := handler.Execute([]string{})

	assert.Nil(t, messages)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrInvalidArguments.Error())
}

func TestUpdateExecute_ReturnsError_WhenIDIsInvalid(t *testing.T) {
	ctl := gomock.NewController(t)
	uCase := usecase.NewMockTaskUseCase(ctl)

	handler := &UpdateCommandHanlder{
		uCase: uCase,
	}

	ID := "123abc"

	messages, err := handler.Execute([]string{"/path/to/file", "update", ID, "new description"})

	assert.Nil(t, messages)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrUnableToParseTaskId.Error())
}

func TestUpdatexecute_ReturnsError_WhenUCaseReturnsError(t *testing.T) {
	ctl := gomock.NewController(t)
	uCase := usecase.NewMockTaskUseCase(ctl)

	ID := uint64(1)
	description := "new description"

	uCaseErr := errors.New("unable to update tasks")

	uCase.EXPECT().UpdateTask(ID, entity.Status(""), description).Return(nil, uCaseErr)

	handler := &UpdateCommandHanlder{
		uCase:  uCase,
	}

	messages, err := handler.Execute([]string{"/path/to/file", "update", strconv.FormatUint(ID, 10), description})

	assert.Nil(t, messages)
	assert.Error(t, err)
	assert.ErrorIs(t, err, uCaseErr)
}

func TestUpdateExecute_ReturnsNils_WhenExecutionIsSuccessful(t *testing.T) {
	ctl := gomock.NewController(t)
	uCase := usecase.NewMockTaskUseCase(ctl)

	description := "new description"
	ID := uint64(1)

	uCase.EXPECT().UpdateTask(ID, entity.Status(""), description).Return(&entity.Task{ID: ID}, nil)

	handler := &UpdateCommandHanlder{
		uCase:  uCase,
	}

	messages, err := handler.Execute([]string{"/path/to/file", "update", strconv.FormatUint(ID, 10), description})

	assert.Nil(t, err)
	assert.Nil(t, messages)
}
