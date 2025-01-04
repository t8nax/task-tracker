package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	storage "github.com/t8nax/task-tracker/internal/storage/fake"
)

func TestNewTaskService_ReturnsInstance_WhenStorageIsValid(t *testing.T) {
	storage := storage.FakeStorage{}
	service := NewTaskService(&storage)

	assert.NotNil(t, service)
}

func TestNewTaskService_Panics_WhenStorageIsNil(t *testing.T) {
	assert.PanicsWithValue(t, storageMustNotBeNilStr, func() {
		NewTaskService(nil)
	})
}
