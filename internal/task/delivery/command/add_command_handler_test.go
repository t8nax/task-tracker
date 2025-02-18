package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/t8nax/task-tracker/internal/task/repository"
	"github.com/t8nax/task-tracker/internal/task/usecase"
)

func TestExecute_ReturnsError_WhenArgLengthIsInvalid(t *testing.T) {
	repo := repository.NewFakeRepository()
	uCase := usecase.NewTaskUseCase(repo)

	handler := &AddCommandHanlder{
		uCase: uCase,
	}

	messages, err := handler.Execute([]string{})
	
	assert.Nil(t, messages)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrInvalidArguments.Error())
}
