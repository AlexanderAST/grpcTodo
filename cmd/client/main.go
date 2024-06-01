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

	flag.Parse()

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := generatedProto.NewUserClient(conn)

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
	}
}
