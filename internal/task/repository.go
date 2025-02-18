package task

import "github.com/t8nax/task-tracker/internal/task/entity"

type Repository interface {
	GetAll() ([]entity.Task, error)
	UpdateAll([]entity.Task) error
}
