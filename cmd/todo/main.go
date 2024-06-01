package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"todoGRPC/pkg/handler"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := handler.Start(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
