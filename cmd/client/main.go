package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"todoGRPC/pkg/generatedProto"
)

func main() {
	serverAddr := flag.String("server", ":8080", "grpc todo")
	id := flag.String("id", "", "User ID")
	email := flag.String("email", "", "User email")
	password := flag.String("password", "", "User password")
	method := flag.String("method", "", "method")
	name := flag.String("name", "", "names")
	description := flag.String("description", "", "description")
	time := flag.String("time", "", "time")
	folderID := flag.String("folderID", "", "folderID")
	userID := flag.String("userID", "", "userID")
	taskID := flag.String("taskID", "", "taskID")

	flag.Parse()

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := generatedProto.NewUserClient(conn)
	clientTask := generatedProto.NewTodoClient(conn)

	switch *method {
	case "signup":
		req := &generatedProto.SignUpRequest{
			Id:       *id,
			Email:    *email,
			Password: *password,
		}

		res, err := client.SignUp(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.Response)
	case "signin":
		req := &generatedProto.SignInRequest{
			Id:       *id,
			Email:    *email,
			Password: *password,
		}
		res, err := client.SignIn(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.Token)

	case "createFolder":
		req := &generatedProto.CreateFolderRequest{
			Id:   *id,
			Name: *name,
		}

		res, err := clientTask.CreateFolder(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.Response)

	case "deleteFolder":
		req := &generatedProto.DeleteFolderRequest{FolderID: *id}

		res, err := clientTask.DeleteFolder(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(res.Response)

	case "getAllInFolder":
		req := &generatedProto.GetAllInFolderRequest{FolderID: *folderID}

		res, err := clientTask.GetAllInFolder(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		for _, i := range res.TaskID {
			fmt.Println(i)
		}

	case "createTask":
		req := &generatedProto.CreateTaskRequest{
			Id:          *id,
			Name:        *name,
			Description: *description,
			Time:        *time,
			UserId:      *userID,
			FolderId:    *folderID,
		}

		res, err := clientTask.CreateTask(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(res.Response)
	case "deleteTask":
		req := &generatedProto.DeleteTaskRequest{Id: *id}

		res, err := clientTask.DeleteTask(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(res.Response)
	case "openTask":
		req := &generatedProto.GetAllOneTaskRequest{Id: *taskID}

		result, err := clientTask.GetAllOneTask(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result.Response)
	case "updateTask":
		req := &generatedProto.UpdateTaskRequest{
			Id:          *id,
			Name:        *name,
			Description: *description,
			Time:        *time,
			UserId:      *userID,
			FolderId:    *folderID,
		}

		res, err := clientTask.UpdateTask(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(res.Response)
	case "getAllTask":
		req := &generatedProto.GetAllTasksRequest{}

		res, err := clientTask.GetAllTasks(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}

		for _, i := range res.TaskID {
			fmt.Println(i)
		}
	}
}
