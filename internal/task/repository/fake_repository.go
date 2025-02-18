package repository

import (
	"github.com/t8nax/task-tracker/internal/task/entity"
)

type FakeRepository struct {
	entities map[uint64]entity.Task
}

func NewFakeRepository() *FakeRepository {
	return &FakeRepository{
		entities: make(map[uint64]entity.Task),
	}
}

func (s *FakeRepository) GetAll() ([]entity.Task, error) {
	res := make([]entity.Task, 0, len(s.entities))
	for _, value := range s.entities {
		res = append(res, value)
	}

	return res, nil
}

func (s *FakeRepository) UpdateAll(tasks []entity.Task) error {
	for id := range s.entities {
		delete(s.entities, id)
	}

	for _, task := range tasks {
		s.entities[task.ID] = task
	}

	return nil
}
