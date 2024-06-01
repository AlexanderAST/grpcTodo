package store

import "todoGRPC/pkg/model"

type UserRepository interface {
	CreateUser(u *model.User) (string, error)
	GetUser(id string) (*model.User, error)
	Login(id string, u *model.User) (string, error)
}
