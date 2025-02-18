package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/t8nax/task-tracker/internal/task/entity"
	"github.com/t8nax/task-tracker/internal/task/repository"
)

func TestUpdate_ReturnsError_WhenErrorOnGettingTasks(t *testing.T) {
	ctl := gomock.NewController(t)
	repo := repository.NewMockRepository(ctl)
	uCase := NewTaskUseCase(repo)

	repoErr := errors.New("DB is down")
	repo.EXPECT().GetAll().Return(nil, repoErr)

	task, err := uCase.UpdateTask(1, "", "")

	assert.Nil(t, task)
	assert.Error(t, err)
	assert.ErrorIs(t, err, errRepoGetTasks)
	assert.Contains(t, err.Error(), repoErr.Error())
}

func TestUpdate_ReturnsError_WhenTaskNotFound(t *testing.T) {
	repo := repository.NewFakeRepository()
	uCase := NewTaskUseCase(repo)

	ID := uint64(1)
	task, err := uCase.UpdateTask(1, "", "")

	assert.Nil(t, task)
	assert.Error(t, err)
	assert.Error(t, getErrTaskNotFound(ID), err.Error())
}

func TestUpdate_ReturnsError_WhenStatusIsInvalid(t *testing.T) {
	repo := repository.NewFakeRepository()
	uCase := NewTaskUseCase(repo)

	ID := uint64(1)
	now := time.Now()

	repo.UpdateAll([]entity.Task{{
		ID:          ID,
		Description: "Description",
		Status:      entity.StatusDone,
		CreatedAt:   now.Add(-24 * time.Hour),
		UpdatedAt:   now.Add(-24 * time.Hour),
	}})

	task, err := uCase.UpdateTask(ID, entity.StatusToDo, "")

	assert.Nil(t, task)
	assert.Error(t, err)
	assert.EqualError(t, getErrInvalidStatus(), err.Error())
}

func TestUpdate_SuccessfulyUpdates_WhenStatusIsValid(t *testing.T) {
	repo := repository.NewFakeRepository()
	uCase := NewTaskUseCase(repo)

	ID := uint64(1)
	status := entity.StatusInProgress
	now := time.Now()

	repo.UpdateAll([]entity.Task{{
		ID:          ID,
		Description: "Description",
		Status:      entity.StatusToDo,
		CreatedAt:   now.Add(-24 * time.Hour),
		UpdatedAt:   now.Add(-24 * time.Hour),
	}})

	task, err := uCase.UpdateTask(ID, status, "")

	assert.Nil(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, status, task.Status)
	assert.WithinDuration(t, now, task.UpdatedAt, time.Second)
}

func TestUpdate_SuccessfulyUpdates_WhenDescriptionIsValid(t *testing.T) {
	repo := repository.NewFakeRepository()
	uCase := NewTaskUseCase(repo)

	now := time.Now()
	ID := uint64(1)

	repo.UpdateAll([]entity.Task{{
		ID:          ID,
		Description: "Old description",
		Status:      entity.StatusToDo,
		CreatedAt:   now.Add(-24 * time.Hour),
		UpdatedAt:   now.Add(-24 * time.Hour),
	}})

	now = time.Now()
	description := "New description"
	task, err := uCase.UpdateTask(1, "", description)

	assert.NoError(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, description, task.Description)
	assert.WithinDuration(t, now, task.UpdatedAt, time.Second)
}

func TestUpdaet_NoUpdates_WhenInputParamsIsEmpty(t *testing.T) {
	repo := repository.NewFakeRepository()
	uCase := NewTaskUseCase(repo)

	time := time.Now().Add(-24 * time.Hour)
	ID := uint64(1)
	description := "Description"
	status := entity.StatusToDo

	repo.UpdateAll([]entity.Task{{
		ID:          ID,
		Description: description,
		Status:      status,
		CreatedAt:   time,
		UpdatedAt:   time,
	}})

	task, err := uCase.UpdateTask(ID, "", "")

	assert.NoError(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, description, task.Description)
	assert.Equal(t, status, task.Status)
	assert.Equal(t, time, task.UpdatedAt)
}

func TestUpdate_ReturnsError_WhenErrorOnUpdatingTask(t *testing.T) {
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

	task, err := uCase.UpdateTask(1, entity.StatusDone, "")

	assert.Nil(t, task)
	assert.Error(t, err)
	assert.ErrorIs(t, err, errRepoUpdateTask)
	assert.Contains(t, err.Error(), repoErr.Error())
}
