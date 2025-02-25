package usecase

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/t8nax/task-tracker/internal/task/entity"
	"github.com/t8nax/task-tracker/internal/task/repository"
)

func TestGetAll_ReturnsError_WhenErrorOnGettingTasks(t *testing.T) {
	ctl := gomock.NewController(t)
	repo := repository.NewMockRepository(ctl)
	uCase := NewTaskUseCase(repo)

	repoErr := errors.New("DB is down")
	repo.EXPECT().GetAll().Return(nil, repoErr)

	tasks, err := uCase.GetAllTasks(entity.StatusNone)

	assert.Nil(t, tasks)
	assert.Error(t, err)
	assert.ErrorIs(t, err, errRepoGetTasks)
	assert.Contains(t, err.Error(), repoErr.Error())
}

func TestGetAll_ReturnsAllTasks_WhenStatusIsNone(t *testing.T) {
	ctl := gomock.NewController(t)
	repo := repository.NewMockRepository(ctl)
	uCase := NewTaskUseCase(repo)

	repoTasks := []entity.Task{{
		ID:     1,
		Status: entity.StatusDone,
	}, {
		ID:     2,
		Status: entity.StatusInProgress,
	}}

	repo.EXPECT().GetAll().Return(repoTasks, nil)

	tasks, err := uCase.GetAllTasks(entity.StatusNone)

	assert.Nil(t, err)
	assert.Equal(t, len(repoTasks), len(tasks))

	for i, task := range tasks {
		assert.Equal(t, task.ID, repoTasks[i].ID)
	}
}

func TestGetAll_ReturnsFilteredTasks_WhenStatusMatch(t *testing.T) {
	ctl := gomock.NewController(t)
	repo := repository.NewMockRepository(ctl)
	uCase := NewTaskUseCase(repo)

	repoTasks := []entity.Task{{
		ID:     1,
		Status: entity.StatusDone,
	}, {
		ID:     2,
		Status: entity.StatusInProgress,
	}}

	repo.EXPECT().GetAll().Return(repoTasks, nil)

	const filteredTaskCount = 1
	filteredTask := repoTasks[0]

	tasks, err := uCase.GetAllTasks(entity.StatusDone)

	assert.Nil(t, err)
	assert.Equal(t, filteredTaskCount, len(tasks))
	assert.Equal(t, filteredTask.ID, tasks[0].ID)
}

func TestGetAll_ReturnsEmptySlice_WhenStatusdDoesNotMatch(t *testing.T) {
	ctl := gomock.NewController(t)
	repo := repository.NewMockRepository(ctl)
	uCase := NewTaskUseCase(repo)

	repoTasks := []entity.Task{{
		ID:     1,
		Status: entity.StatusDone,
	}, {
		ID:     2,
		Status: entity.StatusInProgress,
	}}

	repo.EXPECT().GetAll().Return(repoTasks, nil)

	tasks, err := uCase.GetAllTasks(entity.StatusToDo)

	assert.Nil(t, err)
	assert.NotNil(t, tasks)
	assert.Empty(t, tasks)
}