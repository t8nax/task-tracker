package task

import "github.com/t8nax/task-tracker/internal/task/entity"

type TaskUseCase interface {
	GetAllTasks(status entity.Status) ([]entity.Task, error)
	AddTask(description string) (*entity.Task, error)
	UpdateTask(ID uint64, status entity.Status, description string) (*entity.Task, error)
	DeleteTask(ID uint64) error
}
