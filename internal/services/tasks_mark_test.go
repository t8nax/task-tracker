package tasksrv

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/t8nax/task-tracker/internal/models"
	fake_storage "github.com/t8nax/task-tracker/internal/storage/fake"
	mock_storage "github.com/t8nax/task-tracker/internal/storage/mocks"
)

func TestMarkTask_SuccessfulyMarksTask_WhenInputIsValid(t *testing.T) {
	storage := fake_storage.NewFakeStorage()
	service := NewTaskService(storage)

	now := time.Now()

	storage.UpdateAll([]models.Task{{
		ID:          1,
		Description: "Task 1",
		Status:      models.StatusToDo,
		CreatedAt:   now,
		UpdatedAt:   now.Add(-24 * time.Hour),
	}})

	now = time.Now()
	task, err := service.MarkTask(1, models.StatusDone)

	assert.NoError(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, models.StatusDone, task.Status)
	assert.WithinDuration(t, now, task.UpdatedAt, time.Second)
}

func TestMarkTask_ReturnsError_WhenIDIsInvalid(t *testing.T) {
	storage := fake_storage.NewFakeStorage()
	service := NewTaskService(storage)

	ID := uint64(1)
	task, err := service.MarkTask(ID, models.StatusDone)

	assert.Nil(t, task)
	assert.Error(t, err)
	assert.EqualError(t, getErrTaskNotFound(ID), err.Error())
}

func TestMarkTask_ReturnsError_WhenStatusIsInvalid(t *testing.T) {
	storage := fake_storage.NewFakeStorage()
	service := NewTaskService(storage)

	tests := []struct {
		name   string
		status models.Status
	}{
		{
			name:   "StatusIsRandomStr",
			status: models.Status("invalid_status"),
		},
		{
			name:   "StatusIsNotSuitable",
			status: models.StatusToDo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task, err := service.MarkTask(1, tt.status)

			assert.Nil(t, task)
			assert.Error(t, err)
			assert.EqualError(t, getErrInvalidStatus(), err.Error())
		})
	}
}

func TestMarkTask_ReturnsError_WhenFailedToGetTasks(t *testing.T) {
	ctl := gomock.NewController(t)
	storage := mock_storage.NewMockStorage(ctl)
	service := NewTaskService(storage)

	defer ctl.Finish()

	storageErr := errors.New("DB is down")
	storage.EXPECT().GetAll().Return(nil, storageErr)

	task, err := service.MarkTask(1, models.StatusDone)

	assert.Nil(t, task)
	assert.Error(t, err)
	assert.ErrorIs(t, err, storageErr)
	assert.Contains(t, err.Error(), storageErr.Error())
}

func TestMarkTask_ReturnsError_WhenFailedToUpdateTasks(t *testing.T) {
	ctl := gomock.NewController(t)
	storage := mock_storage.NewMockStorage(ctl)
	service := NewTaskService(storage)

	defer ctl.Finish()

	storage.EXPECT().GetAll().Return([]models.Task{
		{
			ID:     1,
			Status: models.StatusToDo,
		},
	}, nil)

	storageErr := errors.New("DB is down")
	storage.EXPECT().UpdateAll(gomock.Any()).Return(storageErr)

	task, err := service.MarkTask(1, models.StatusDone)

	assert.Nil(t, task)
	assert.Error(t, err)
	assert.ErrorIs(t, err, storageErr)
	assert.Contains(t, err.Error(), storageErr.Error())
}
