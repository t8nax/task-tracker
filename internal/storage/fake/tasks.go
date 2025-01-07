package storage

import "github.com/t8nax/task-tracker/internal/models"

type FakeStorage struct {
	entities map[uint64]models.Task
}

func NewFakeStorage() *FakeStorage {
	return &FakeStorage{
		entities: make(map[uint64]models.Task),
	}
}

func (s *FakeStorage) GetAll() ([]models.Task, error) {
	res := make([]models.Task, 0, len(s.entities))
	for _, value := range s.entities {
		res = append(res, value)
	}

	return res, nil
}

func (s *FakeStorage) UpdateAll(tasks []models.Task) error {
	for id := range s.entities {
		delete(s.entities, id)
	}

	for _, task := range tasks {
		s.entities[task.ID] = task
	}

	return nil
}
