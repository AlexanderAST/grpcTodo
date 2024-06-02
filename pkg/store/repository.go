package store

import "todoGRPC/pkg/model"

type UserRepository interface {
	CreateUser(u *model.User) (string, error)
	GetUser(id string) (*model.User, error)
	Login(id string, u *model.User) (string, error)
}

type TaskRepository interface {
	CreateFolder(f *model.Folder) (string, error)
	DeleteFolder(id string) string
	GetAllInFolder(id string) ([]int, error)
	CreateTask(t *model.Task) (string, error)
	DeleteTask(id string) string
	GetOpenTask(id string) (string, error)
	GetAllTasks() ([]int, error)
	UpdateTask(id string, t *model.Task) (string, error)
}
