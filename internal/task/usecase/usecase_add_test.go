package usecase

import (
	"errors"
	"math"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/t8nax/task-tracker/internal/task/entity"
	"github.com/t8nax/task-tracker/internal/task/repository"
	mathutils "github.com/t8nax/task-tracker/pkg/math"
)

func TestAddTask_SuccessfullyAddsTask_WhenDescriptionIsVaild(t *testing.T) {
	repo := repository.NewFakeRepository()
	uCase := NewTaskUseCase(repo)

	now := time.Now()

	repo.UpdateAll([]entity.Task{{
		ID:          1,
		Description: "Task 1",
		Status:      entity.StatusToDo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}})

	in := "Task 2"
	task, err := uCase.AddTask(in)

	assert.NoError(t, err)
	assert.Equal(t, in, task.Description)
	assert.Equal(t, entity.StatusToDo, task.Status)
	assert.Equal(t, uint64(2), task.ID)
	assert.WithinDuration(t, now, task.CreatedAt, time.Second)
	assert.WithinDuration(t, now, task.UpdatedAt, time.Second)
}

func TestAddTask_ReturnsError_WhenDescriptionIsEmpty(t *testing.T) {
	repo := repository.NewFakeRepository()
	uCase := NewTaskUseCase(repo)

	in := ""
	task, err := uCase.AddTask(in)

	assert.Nil(t, task)
	assert.Error(t, err)
	assert.EqualError(t, err, errEmptyDescription.Error())
}

func TestAddTask_ReturnsError_WhenStorageGetAllFails(t *testing.T) {
	ctl := gomock.NewController(t)
	repo := repository.NewMockRepository(ctl)
	uCase := NewTaskUseCase(repo)
	defer ctl.Finish()

	repoErr := errors.New("DB is down")
	repo.EXPECT().GetAll().Return(nil, repoErr)

	in := "description"
	task, err := uCase.AddTask(in)

	assert.Nil(t, task)
	assert.Error(t, err)
	assert.ErrorIs(t, err, errRepoGetTasks)
	assert.Contains(t, err.Error(), repoErr.Error())
}

func TestAddTask_ReturnsError_WhenStorageUpdateAllFails(t *testing.T) {
	ctl := gomock.NewController(t)
	repo := repository.NewMockRepository(ctl)
	uCase := NewTaskUseCase(repo)
	defer ctl.Finish()

	repoErr := errors.New("DB is down")

	repo.EXPECT().GetAll().Return([]entity.Task{}, nil)
	repo.EXPECT().UpdateAll(gomock.Any()).Return(repoErr)

	in := "description"
	task, err := uCase.AddTask(in)

	assert.Nil(t, task)
	assert.Error(t, err)
	assert.ErrorIs(t, err, errRepoAddTask)
	assert.Contains(t, err.Error(), repoErr.Error())
}

func TestAddTasks_ReturnsError_WhenFailedToGenerateTaskId(t *testing.T) {
	ctl := gomock.NewController(t)
	repo := repository.NewMockRepository(ctl)
	uCase := NewTaskUseCase(repo)
	defer ctl.Finish()

	repo.EXPECT().GetAll().Return([]entity.Task{
		{
			ID: math.MaxUint64,
		},
	}, nil)

	in := "description"
	task, err := uCase.AddTask(in)

	assert.Nil(t, task)
	assert.Error(t, err)
	assert.ErrorIs(t, err, errGenerateTaskID)
	assert.Contains(t, err.Error(), mathutils.ErrMaxValue.Error())
}
