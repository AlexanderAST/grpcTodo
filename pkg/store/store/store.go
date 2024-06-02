package store

import (
	"context"
	"github.com/go-redis/redis/v8"
	"todoGRPC/pkg/store"
)

var ctx = context.Background()

type Store struct {
	db             *redis.Client
	userRepository *UserRepository
	taskRepository *TaskRepository
}

func New(db *redis.Client) *Store {
	return &Store{db: db}
}

func (s *Store) User() store.UserRepository {

	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{store: s}

	return s.userRepository
}

func (s *Store) Task() store.TaskRepository {

	if s.taskRepository != nil {
		return s.taskRepository
	}

	s.taskRepository = &TaskRepository{store: s}

	return s.taskRepository
}
