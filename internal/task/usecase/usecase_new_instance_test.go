package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/t8nax/task-tracker/internal/task/repository"
)

func TestNewTaskUseCase_ReturnsInstance_WhenRepositoryIsValid(t *testing.T) {
	repo := repository.NewFakeRepository()
	uCase := NewTaskUseCase(repo)

	assert.NotNil(t, uCase)
}

func TestNewTaskUseCase_Panics_WhenRepositoryIsNil(t *testing.T) {
	assert.PanicsWithValue(t, repoMustNotBeNilStr, func() {
		NewTaskUseCase(nil)
	})
}
