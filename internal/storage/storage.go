package storage

import "github.com/t8nax/task-tracker/internal/models"

type Storage interface {
	ReadAll() ([]*models.Task, error)
}
