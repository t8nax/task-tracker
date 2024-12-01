package jsonstorage

import (
	"encoding/json"
	"fmt"
	"github.com/t8nax/task-tracker/internal/models"
	"github.com/t8nax/task-tracker/pkg/files"
	"os"
)

const fileName = "tasks.json"

type JsonStorage struct{}

func (s *JsonStorage) ReadAll() ([]*models.Task, error) {
	if !files.Exists(fileName) {
		return []*models.Task{}, nil
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("unable to open database file %s: %w", fileName, err)
	}

	defer closeFile(file)

	tasks := make([]*models.Task, 0)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)

	if err != nil {
		return nil, fmt.Errorf("failed to decode json database file %s: %w", fileName, err)
	}

	return tasks, nil
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Printf("WARNING: Unable to close database file %s: %s\n", fileName, err)
	}
}
