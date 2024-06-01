package handler

import (
	"context"
	"log"
	"todoGRPC/pkg/generatedProto"
	"todoGRPC/pkg/model"
)

func (s *server) SignUp(ctx context.Context, req *generatedProto.SignUpRequest) (*generatedProto.SignUpResponse, error) {

	u := &model.User{
		ID:       req.Id,
		Email:    req.Email,
		Password: req.Password,
	}

	result, err := s.store.User().CreateUser(u)
	if err != nil {
		log.Fatal(err)
	}

	return &generatedProto.SignUpResponse{Response: result}, nil
}

func (s *server) SignIn(ctx context.Context, req *generatedProto.SignInRequest) (*generatedProto.SignInResponse, error) {

	u := &model.User{
		ID:       req.Id,
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := s.store.User().GetUser(u.ID)
	if err != nil {
		log.Fatal(err)
	}

	userResult, err := s.store.User().Login(user.ID, u)
	if err != nil {
		log.Fatal(err)
	}

	return &generatedProto.SignInResponse{Token: userResult}, nil
}
