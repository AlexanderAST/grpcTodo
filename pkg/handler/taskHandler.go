package handler

import (
	"context"
	"log"
	"todoGRPC/pkg/generatedProto"
	"todoGRPC/pkg/model"
)

func (s *server) CreateFolder(ctx context.Context, req *generatedProto.CreateFolderRequest) (*generatedProto.CreateFolderResponse, error) {

	f := &model.Folder{
		ID:   req.Id,
		Name: req.Name,
	}

	result, err := s.store.Task().CreateFolder(f)
	if err != nil {
		log.Fatal(err)
	}

	return &generatedProto.CreateFolderResponse{Response: result}, nil
}

func (s *server) DeleteFolder(ctx context.Context, req *generatedProto.DeleteFolderRequest) (*generatedProto.DeleteFolderResponse, error) {

	result := s.store.Task().DeleteFolder(req.FolderID)

	return &generatedProto.DeleteFolderResponse{Response: result}, nil
}

func (s *server) GetAllInFolder(ctx context.Context, req *generatedProto.GetAllInFolderRequest) (*generatedProto.GetAllInFolderResponse, error) {
	result, err := s.store.Task().GetAllInFolder(req.FolderID)
	if err != nil {
		log.Fatal(err)
	}

	var int64s []int64
	for _, i := range result {
		int64s = append(int64s, int64(i))
	}

	return &generatedProto.GetAllInFolderResponse{TaskID: int64s}, nil
}

func (s *server) CreateTask(ctx context.Context, req *generatedProto.CreateTaskRequest) (*generatedProto.CreateTaskResponse, error) {

	t := &model.Task{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Time:        req.Time,
		UserID:      req.UserId,
		FolderID:    req.FolderId,
	}

	res, err := s.store.Task().CreateTask(t)
	if err != nil {
		log.Fatal(err)
	}

	return &generatedProto.CreateTaskResponse{Response: res}, nil
}

func (s *server) DeleteTask(ctx context.Context, req *generatedProto.DeleteTaskRequest) (*generatedProto.DeleteTaskResponse, error) {
	result := s.store.Task().DeleteTask(req.Id)

	return &generatedProto.DeleteTaskResponse{Response: result}, nil
}

func (s *server) GetAllOneTask(ctx context.Context, req *generatedProto.GetAllOneTaskRequest) (*generatedProto.GetAllOneTaskResponse, error) {
	result, err := s.store.Task().GetOpenTask(req.Id)
	if err != nil {
		return nil, err
	}

	return &generatedProto.GetAllOneTaskResponse{Response: result}, nil

}

func (s *server) GetAllTasks(ctx context.Context, req *generatedProto.GetAllTasksRequest) (*generatedProto.GetAllTasksResponse, error) {
	result, err := s.store.Task().GetAllTasks()
	if err != nil {
		return nil, err
	}

	var int64s []int64
	for _, i := range result {
		int64s = append(int64s, int64(i))
	}

	return &generatedProto.GetAllTasksResponse{TaskID: int64s}, nil
}

func (s *server) UpdateTask(ctx context.Context, req *generatedProto.UpdateTaskRequest) (*generatedProto.UpdateTaskResponse, error) {
	t := &model.Task{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Time:        req.Time,
		UserID:      req.UserId,
		FolderID:    req.FolderId,
	}

	res, err := s.store.Task().UpdateTask(req.Id, t)
	if err != nil {
		log.Fatal(err)
	}

	return &generatedProto.UpdateTaskResponse{Response: res}, nil
}
