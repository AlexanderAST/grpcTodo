package handler

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"todoGRPC/pkg/generatedProto"
	"todoGRPC/pkg/store"
	storeDb "todoGRPC/pkg/store/store"
)

type server struct {
	generatedProto.UnimplementedUserServer
	generatedProto.UnimplementedTodoServer
	store store.Store
}

func newServer(store store.Store) *server {
	return &server{
		store: store,
	}
}

func Start() error {
	rdb, err := newRedisDB()
	if err != nil {
		log.Fatal(err)
	}
	defer rdb.Close()

	s := grpc.NewServer()
	store := storeDb.New(rdb)
	srv := newServer(store)

	generatedProto.RegisterUserServer(s, srv)
	generatedProto.RegisterTodoServer(s, srv)

	port := viper.GetString("port")

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("server started")

	return s.Serve(l)
}

func newRedisDB() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("rdb.address"),
		Password: os.Getenv("DBPASSWORD"),
		DB:       viper.GetInt("rdb.db"),
	})

	rdb.Ping(context.Background())

	return rdb, nil
}
