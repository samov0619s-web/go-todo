package storage

import (
	"errors"
	"go-todo/internal/model"
	"sync"
)

type MemoryStorage struct {
	mu     sync.Mutex
	todos  []model.Todo
	nextID int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		todos:  []model.Todo{},
		nextID: 1,
	}
}

func (s *MemoryStorage) GetAll() []model.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.todos
}

func (s *MemoryStorage) Create(title string) model.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo := model.Todo{
		ID:    s.nextID,
		Title: title,
		Done:  false,
	}

	s.nextID++
	s.todos = append(s.todos, todo)

	return todo
}

func (s *MemoryStorage) Update(id int, title *string, done *bool) (model.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, t := range s.todos {
		if t.ID == id {
			if title != nil {
				s.todos[i].Title = *title
			}
			if done != nil {
				s.todos[i].Done = *done
			}
			return s.todos[i], nil
		}
	}

	return model.Todo{}, errors.New("not found")
}

func (s *MemoryStorage) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, t := range s.todos {
		if t.ID == id {
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			return nil
		}
	}

	return errors.New("not found")
}
