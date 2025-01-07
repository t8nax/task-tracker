package tasksrv

import (
	"testing"

	"github.com/stretchr/testify/assert"
	fake_storage "github.com/t8nax/task-tracker/internal/storage/fake"
)

func TestNewTaskService_ReturnsInstance_WhenStorageIsValid(t *testing.T) {
	storage := fake_storage.NewFakeStorage()
	service := NewTaskService(storage)

	assert.NotNil(t, service)
}

func TestNewTaskService_Panics_WhenStorageIsNil(t *testing.T) {
	assert.PanicsWithValue(t, storageMustNotBeNilStr, func() {
		NewTaskService(nil)
	})
}
