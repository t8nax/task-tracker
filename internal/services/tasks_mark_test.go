package services

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/t8nax/task-tracker/internal/models"
	storage "github.com/t8nax/task-tracker/internal/storage/fake"
	mock_storage "github.com/t8nax/task-tracker/internal/storage/mocks"
)

func TestMark_ReturnsNil_WhenInputIsValid(t *testing.T) {
	storage := storage.FakeStorage{}
	service := &TaskService{
		storage: &storage,
	}

	now := time.Now()

	storage.UpdateAll([]models.Task{{
		ID:          1,
		Description: "Task 1",
		Status:      models.StatusToDo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}})

	err := service.Mark(1, models.StatusDone)

	assert.NoError(t, err)
}

func TestMark_ReturnsError_WhenIDIsInvalid(t *testing.T) {
	storage := storage.FakeStorage{}
	service := &TaskService{
		storage: &storage,
	}

	ID := uint64(0)
	err := service.Mark(ID, models.StatusDone)

	assert.Error(t, err)
	assert.EqualError(t, getErrTaskNotFound(ID), err.Error())
}

func TestMark_ReturnsError_WhenStatusIsInvalid(t *testing.T) {
	storage := storage.FakeStorage{}
	service := &TaskService{
		storage: &storage,
	}

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
			err := service.Mark(1, tt.status)

			assert.Error(t, err)
			assert.EqualError(t, getErrInvalidStatus(), err.Error())
		})
	}
}

func TestAddTasks_ReturnsError_WhenFailedToGetTasks(t *testing.T) {
	ctl := gomock.NewController(t)
	storage := mock_storage.NewMockStorage(ctl)
	service := &TaskService{
		storage: storage,
	}
	defer ctl.Finish()

	storageErr := errors.New("DB is down")
	storage.EXPECT().GetAll().Return(nil, storageErr)

	err := service.Mark(1, models.StatusDone)

	assert.Error(t, err)
	assert.ErrorIs(t, err, storageErr)
	assert.Contains(t, err.Error(), storageErr.Error())
}

func TestAddTasks_ReturnsError_WhenFailedToUpdateTasks(t *testing.T) {
	ctl := gomock.NewController(t)
	storage := mock_storage.NewMockStorage(ctl)
	service := &TaskService{
		storage: storage,
	}
	defer ctl.Finish()

	storage.EXPECT().GetAll().Return([]models.Task{
		{
			ID:     1,
			Status: models.StatusToDo,
		},
	}, nil)

	storageErr := errors.New("DB is down")
	storage.EXPECT().UpdateAll(gomock.Any()).Return(storageErr)

	err := service.Mark(1, models.StatusDone)

	assert.Error(t, err)
	assert.ErrorIs(t, err, storageErr)
	assert.Contains(t, err.Error(), storageErr.Error())
}
