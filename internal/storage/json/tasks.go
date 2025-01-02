package storage

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"

	"github.com/t8nax/task-tracker/internal/models"
	"github.com/t8nax/task-tracker/pkg/files"
)

const fileName = "tasks.json"

type JsonStorage struct{}

func (s *JsonStorage) GetAll() ([]models.Task, error) {
	if !files.Exists(fileName) {
		return []models.Task{}, nil
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("unable to open database file %s: %w", fileName, err)
	}

	defer closeFile(file)
	return decode(file)
}

func decode(file *os.File) ([]models.Task, error) {
	tasks := make([]models.Task, 0)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&tasks)

	if err != nil {
		return nil, fmt.Errorf("failed to decode json database file %s: %w", fileName, err)
	}

	return tasks, nil
}

func (s *JsonStorage) UpdateAll(tasks []models.Task) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, fs.FileMode(0644))

	if err != nil {
		return fmt.Errorf("unable to open database file %s: %w", fileName, err)
	}

	defer closeFile(file)
	bytes, err := json.Marshal(&tasks)
	if err != nil {
		return fmt.Errorf("unable to encode tasks to json: %w", err)
	}

	_, err = file.Write(bytes)

	if err != nil {
		return fmt.Errorf("unable to write tasks to database file: %w", err)
	}

	return nil
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Printf("WARNING: Unable to close database file %s: %s\n", fileName, err)
	}
}
