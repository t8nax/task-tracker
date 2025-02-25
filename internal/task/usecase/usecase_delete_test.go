package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	entity "github.com/t8nax/task-tracker/internal/task/entity"
	"github.com/t8nax/task-tracker/internal/task/repository"
)

func TestDelete_ReturnsError_WhenErrorOnGettingTasks(t *testing.T) {
	ctl := gomock.NewController(t)
	repo := repository.NewMockRepository(ctl)
	uCase := NewTaskUseCase(repo)

	repoErr := errors.New("DB is down")
	repo.EXPECT().GetAll().Return(nil, repoErr)

	err := uCase.DeleteTask(1)

	assert.Error(t, err)
	assert.ErrorIs(t, err, errRepoGetTasks)
	assert.Contains(t, err.Error(), repoErr.Error())
}

func TestDelete_ReturnsError_WhenTaskNotFound(t *testing.T) {
	repo := repository.NewFakeRepository()
	uCase := NewTaskUseCase(repo)

	ID := uint64(1)
	err := uCase.DeleteTask(1)

	assert.Error(t, err)
	assert.Error(t, getErrTaskNotFound(ID), err.Error())
}


func TestDelete_ReturnsError_WhenErrorOnDeletingTask(t *testing.T) {
	ctl := gomock.NewController(t)
	repo := repository.NewMockRepository(ctl)
	uCase := NewTaskUseCase(repo)

	ID := uint64(1)
	time := time.Now().Add(-24 * time.Hour)

	repo.EXPECT().GetAll().Return([]entity.Task{{
		ID:          ID,
		Description: "Description",
		Status:      entity.StatusToDo,
		CreatedAt:   time,
		UpdatedAt:   time,
	}}, nil)

	repoErr := errors.New("DB is down")
	repo.EXPECT().UpdateAll(gomock.Any()).Return(repoErr)

	err := uCase.DeleteTask(ID)

	assert.Error(t, err)
	assert.ErrorIs(t, err, errRepoDeleteTask)
	assert.Contains(t, err.Error(), repoErr.Error())
}

func TestUpdate_SuccessfulyDeletion(t *testing.T) {
	repo := repository.NewFakeRepository()
	uCase := NewTaskUseCase(repo)

	now := time.Now()
	ID := uint64(1)

	repo.UpdateAll([]entity.Task{{
		ID:          ID,
		Description: "Description",
		Status:      entity.StatusToDo,
		CreatedAt:   now.Add(-24 * time.Hour),
		UpdatedAt:   now.Add(-24 * time.Hour),
	}})

	err := uCase.DeleteTask(1)

	assert.Nil(t, err)
}
