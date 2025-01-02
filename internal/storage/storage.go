package storage

import "github.com/t8nax/task-tracker/internal/models"

type Storage interface {
	GetAll() ([]models.Task, error)
	UpdateAll([]models.Task) error
}
