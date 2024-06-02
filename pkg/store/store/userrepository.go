package store

import (
	"encoding/json"
	"todoGRPC/pkg/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) CreateUser(u *model.User) (string, error) {
	jsonData, err := json.Marshal(u)
	if err != nil {
		return "", err
	}

	err = r.store.db.Set(ctx, "user:"+u.ID, jsonData, 0).Err()
	if err != nil {
		return "", err
	}
	return "success", nil
}

func (r *UserRepository) GetUser(id string) (*model.User, error) {
	jsonData, err := r.store.db.Get(ctx, "user:"+id).Result()
	if err != nil {
		return nil, err
	}

	var user model.User
	err = json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Login(id string, u *model.User) (string, error) {
	user, err := r.GetUser(id)
	if err != nil {
		return "", err
	}

	if user.Email != u.Email || user.Password != u.Password {
		return "incorrect email or password", nil
	}

	return "login success", nil
}
