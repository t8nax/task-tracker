package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/t8nax/task-tracker/internal/models"
	"github.com/t8nax/task-tracker/internal/storage"
	mathutils "github.com/t8nax/task-tracker/pkg/math"
)

var ErrEmptyDescription = errors.New("description must not be empty")
var ErrStorageGetTasks = errors.New("failed to get tasks from storage")
var ErrStorageAddTask = errors.New("failed to add task to storage")
var ErrGenerateTaskID = errors.New("failed to generate task ID")

var StorageMustNotBeNilStr = "storage must not be nil"

type TaskService struct {
	storage storage.Storage
}

func NewTaskService(s storage.Storage) *TaskService {
	if s == nil {
		panic(StorageMustNotBeNilStr)
	}

	return &TaskService{storage: s}
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	tasks, err := s.storage.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrStorageGetTasks, err)
	}

	return tasks, nil
}

func (s *TaskService) AddTask(description string) (*models.Task, error) {
	if description == "" {
		return nil, ErrEmptyDescription
	}

	tasks, err := s.GetAllTasks()

	if err != nil {
		return nil, err
	}

	now := time.Now()
	ids := make([]uint64, len(tasks))

	for _, task := range tasks {
		ids = append(ids, task.Id)
	}

	id, err := mathutils.GenerateNextNumber(ids)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGenerateTaskID, err)
	}

	task := models.Task{
		Id:          id,
		Description: description,
		Status:      models.StatusToDo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks = append(tasks, task)

	err = s.storage.UpdateAll(tasks)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrStorageAddTask, err)
	}

	return &task, nil
}
