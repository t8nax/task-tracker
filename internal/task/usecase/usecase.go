package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/t8nax/task-tracker/internal/task"
	"github.com/t8nax/task-tracker/internal/task/entity"
	mathutils "github.com/t8nax/task-tracker/pkg/math"
)

var errEmptyDescription = errors.New("description must not be empty")
var errRepoGetTasks = errors.New("failed to get tasks")
var errRepoAddTask = errors.New("failed to add task ")
var errRepoUpdateTask = errors.New("failed to update task")
var errGenerateTaskID = errors.New("failed to generate task ID")

var repoMustNotBeNilStr = "repository must not be nil"

func getErrTaskNotFound(taskID uint64) error {
	return fmt.Errorf("task %d not found", taskID)
}

func getErrInvalidStatus() error {
	return fmt.Errorf("marked status must be \"%s\" or \"%s\"", entity.StatusInProgress, entity.StatusToDo)
}

type TaskUseCase struct {
	repo task.Repository
}

func NewTaskUseCase(r task.Repository) *TaskUseCase {
	if r == nil {
		panic(repoMustNotBeNilStr)
	}

	return &TaskUseCase{
		repo: r,
	}
}

func (u *TaskUseCase) GetAllTasks() ([]entity.Task, error) {
	tasks, err := u.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errRepoGetTasks, err)
	}

	return tasks, nil
}

func (u *TaskUseCase) AddTask(description string) (*entity.Task, error) {
	if description == "" {
		return nil, errEmptyDescription
	}

	tasks, err := u.GetAllTasks()

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

	task := entity.Task{
		ID:          ID,
		Description: description,
		Status:      entity.StatusToDo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks = append(tasks, task)

	err = u.repo.UpdateAll(tasks)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", errRepoAddTask, err)
	}

	return &task, nil
}

func (u *TaskUseCase) UpdateTask(ID uint64, status entity.Status, description string) (*entity.Task, error) {
	tasks, err := u.GetAllTasks()

	if err != nil {
		return nil, err
	}

	var task *entity.Task = nil

	for i := range tasks {
		if tasks[i].ID == ID {
			task = &tasks[i]
		}
	}

	if task == nil {
		return nil, getErrTaskNotFound(ID)
	}

	if status == "" && description == "" {
		return task, nil
	}

	if status != "" {
		if status != entity.StatusDone && status != entity.StatusInProgress {
			return nil, getErrInvalidStatus()
		}

		task.Status = status
	}

	if description != "" {
		task.Description = description
	}
	
	task.UpdatedAt = time.Now()
	err = u.repo.UpdateAll(tasks)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", errRepoUpdateTask, err)
	}

	return task, nil
}