package tasksrv

import (
	"errors"
	"fmt"
	"time"

	"github.com/t8nax/task-tracker/internal/models"
	"github.com/t8nax/task-tracker/internal/storage"
	mathutils "github.com/t8nax/task-tracker/pkg/math"
)

var errEmptyDescription = errors.New("description must not be empty")
var errStorageGetTasks = errors.New("failed to get tasks from storage")
var errStorageAddTask = errors.New("failed to add task to storage")
var errGenerateTaskID = errors.New("failed to generate task ID")

var storageMustNotBeNilStr = "storage must not be nil"

func getErrTaskNotFound(taskID uint64) error {
	return fmt.Errorf("task %d not found", taskID)
}

func getErrInvalidStatus() error {
	return fmt.Errorf("marked status must be \"%s\" or \"%s\"", models.StatusInProgress, models.StatusToDo)
}

type TaskService struct {
	storage storage.Storage
}

func NewTaskService(s storage.Storage) *TaskService {
	if s == nil {
		panic(storageMustNotBeNilStr)
	}

	return &TaskService{
		storage: s,
	}
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	tasks, err := s.storage.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errStorageGetTasks, err)
	}

	return tasks, nil
}

func (s *TaskService) AddTask(description string) (*models.Task, error) {
	if description == "" {
		return nil, errEmptyDescription
	}

	tasks, err := s.GetAllTasks()

	if err != nil {
		return nil, err
	}

	now := time.Now()
	ids := make([]uint64, len(tasks))

	for _, task := range tasks {
		ids = append(ids, task.ID)
	}

	ID, err := mathutils.GenerateNextNumber(ids)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", errGenerateTaskID, err)
	}

	task := models.Task{
		ID:          ID,
		Description: description,
		Status:      models.StatusToDo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks = append(tasks, task)

	err = s.storage.UpdateAll(tasks)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", errStorageAddTask, err)
	}

	return &task, nil
}

func (s *TaskService) MarkTask(ID uint64, status models.Status) (*models.Task, error) {
	if status != models.StatusInProgress && status != models.StatusDone {
		return nil, getErrInvalidStatus()
	}

	task, err := s.update(ID, "", status)

	if err != nil {
		return nil, err
	}

	return task, err
}

func (s *TaskService) UpdateDescription(ID uint64, description string) (*models.Task, error) {
	if description == "" {
		return nil, errEmptyDescription
	}

	task, err := s.update(ID, description, "")

	if err != nil {
		return nil, err
	}

	return task, err
}

func (s *TaskService) update(ID uint64, description string, status models.Status) (*models.Task, error) {
	tasks, err := s.GetAllTasks()

	if err != nil {
		return nil, err
	}

	task, err := s.get(ID, tasks)

	if err != nil {
		return nil, err
	}

	if status != "" {
		task.Status = status
	}
	if description != "" {
		task.Description = description
	}

	task.UpdatedAt = time.Now()
	fmt.Println(task)
	err = s.storage.UpdateAll(tasks)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) get(ID uint64, tasks []models.Task) (*models.Task, error) {
	for i, task := range tasks {
		if task.ID == ID {
			return &tasks[i], nil
		}
	}

	return nil, getErrTaskNotFound(ID)
}
