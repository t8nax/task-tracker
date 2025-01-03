package services

import (
	"errors"
	"math"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/t8nax/task-tracker/internal/models"
	storage "github.com/t8nax/task-tracker/internal/storage/fake"
	mock_storage "github.com/t8nax/task-tracker/internal/storage/mocks"
	mathutils "github.com/t8nax/task-tracker/pkg/math"
)

func TestAddTask_SuccessfullyAddsTask_WhenDescriptionIsVaild(t *testing.T) {
	storage := storage.FakeStorage{}
	service := &TaskService{
		storage: &storage,
	}

	now := time.Now()

	storage.UpdateAll([]models.Task{{
		Id:          1,
		Description: "Task 1",
		Status:      models.StatusToDo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}})

	in := "Task 2"
	task, err := service.AddTask(in)

	assert.NoError(t, err)
	assert.Equal(t, in, task.Description)
	assert.Equal(t, models.StatusToDo, task.Status)
	assert.Equal(t, uint64(2), task.Id)
	assert.WithinDuration(t, now, task.CreatedAt, time.Second)
	assert.WithinDuration(t, now, task.UpdatedAt, time.Second)
}

func TestAddTask_ReturnsError_WhenDescriptionIsEmpty(t *testing.T) {
	storage := storage.FakeStorage{}
	service := &TaskService{
		storage: &storage,
	}

	in := ""
	task, err := service.AddTask(in)

	assert.Nil(t, task)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrEmptyDescription.Error())
}

func TestAddTask_ReturnsError_WhenStorageGetAllFails(t *testing.T) {
	ctl := gomock.NewController(t)
	storage := mock_storage.NewMockStorage(ctl)
	service := &TaskService{
		storage: storage,
	}
	defer ctl.Finish()

	storageErr := errors.New("DB is down")
	storage.EXPECT().GetAll().Return(nil, storageErr)

	in := "description"
	task, err := service.AddTask(in)

	assert.Nil(t, task)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrStorageGetTasks)
	assert.Contains(t, err.Error(), storageErr.Error())
}

func TestAddTask_ReturnsError_WhenStorageUpdateAllFails(t *testing.T) {
	ctl := gomock.NewController(t)
	storage := mock_storage.NewMockStorage(ctl)
	service := &TaskService{
		storage: storage,
	}
	defer ctl.Finish()

	storageErr := errors.New("DB is down")

	storage.EXPECT().GetAll().Return([]models.Task{}, nil)
	storage.EXPECT().UpdateAll(gomock.Any()).Return(storageErr)

	in := "description"
	task, err := service.AddTask(in)

	assert.Nil(t, task)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrStorageAddTask)
	assert.Contains(t, err.Error(), storageErr.Error())
}

func TestAddTasks_ReturnsError_WhenFailedToGenerateTaskId(t *testing.T) {
	ctl := gomock.NewController(t)
	storage := mock_storage.NewMockStorage(ctl)
	service := &TaskService{
		storage: storage,
	}
	defer ctl.Finish()

	storage.EXPECT().GetAll().Return([]models.Task{
		{
			Id: math.MaxUint64,
		},
	}, nil)

	in := "description"
	task, err := service.AddTask(in)

	assert.Nil(t, task)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrGenerateTaskID)
	assert.Contains(t, err.Error(), mathutils.ErrMaxValue.Error())
}
