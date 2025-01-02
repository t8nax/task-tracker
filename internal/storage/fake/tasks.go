package storage

import "github.com/t8nax/task-tracker/internal/models"

type FakeStorage struct{}

var entities = map[uint64]models.Task{}

func (s *FakeStorage) GetAll() ([]models.Task, error) {
	res := make([]models.Task, 0, len(entities))
	for _, value := range entities {
		res = append(res, value)
	}

	return res, nil
}

func (s *FakeStorage) UpdateAll(tasks []models.Task) error {
	for id := range entities {
		delete(entities, id)
	}

	for _, task := range tasks {
		entities[task.Id] = task
	}

	return nil
}
