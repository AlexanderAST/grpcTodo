package handler

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"todoGRPC/pkg/generatedProto"
)

type server struct {
	generatedProto.UnimplementedUserServer
}

func Start() error {

	s := grpc.NewServer()

	srv := &server{}

	generatedProto.RegisterUserServer(s, srv)

	l, err := net.Listen("tcp", viper.GetString("port"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("server started")

	return s.Serve(l)
}
