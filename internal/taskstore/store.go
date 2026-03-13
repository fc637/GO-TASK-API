package store

import (
	"sync"
	"github.com/fc637/go-task-api/internal/models"
)

type TaskStore struct {
	mu    sync.Mutex
	tasks []models.Task
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks: []models.Task{},
	}
}
 func (s *TaskStore) Add(task models.Task) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.tasks = append(s.tasks, task)
}

func (s *TaskStore) List() []models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.tasks
}
