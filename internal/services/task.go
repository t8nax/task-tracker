package services

import (
	"github.com/t8nax/task-tracker/internal/models"
	"github.com/t8nax/task-tracker/internal/storage"
)

type TaskService struct {
	storage storage.Storage
}

func NewTaskService(s storage.Storage) *TaskService {
	return &TaskService{storage: s}
}

func (s *TaskService) GetAllTasks() ([]*models.Task, error) {
	return s.storage.ReadAll()
}
