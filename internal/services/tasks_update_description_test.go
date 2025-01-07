package tasksrv

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/t8nax/task-tracker/internal/models"
	fake_storage "github.com/t8nax/task-tracker/internal/storage/fake"
)

func TestUpdateDescription_SuccessfulyUpdates_WhenInputIsValid(t *testing.T) {
	storage := fake_storage.NewFakeStorage()
	service := NewTaskService(storage)

	now := time.Now()

	storage.UpdateAll([]models.Task{{
		ID:          1,
		Description: "Old description",
		Status:      models.StatusToDo,
		CreatedAt:   now,
		UpdatedAt:   now.Add(-24 * time.Hour),
	}})

	now = time.Now()
	description := "New description"
	task, err := service.UpdateDescription(1, description)

	assert.NoError(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, description, task.Description)
	assert.WithinDuration(t, now, task.UpdatedAt, time.Second)
}

func TestUpdateDescription_ReturnsError_WhenIDIsInvalid(t *testing.T) {
	storage := fake_storage.NewFakeStorage()
	service := NewTaskService(storage)

	ID := uint64(1)
	task, err := service.UpdateDescription(ID, "New description")

	assert.Nil(t, task)
	assert.Error(t, err)
	assert.EqualError(t, getErrTaskNotFound(ID), err.Error())
}

func TestUpdateDescription_ReturnsError_WhenDescriptionIsEmptty(t *testing.T) {
	storage := fake_storage.NewFakeStorage()
	service := NewTaskService(storage)

	task, err := service.UpdateDescription(1, "")

	assert.Nil(t, task)
	assert.Error(t, err)
	assert.EqualError(t, errEmptyDescription, err.Error())
}
