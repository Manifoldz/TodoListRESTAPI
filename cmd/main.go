package main

import (
	"log"

	"github.com/Manifoldz/TodoListRESTAPI/internal/server"
	"github.com/Manifoldz/TodoListRESTAPI/pkg/handler"
	"github.com/Manifoldz/TodoListRESTAPI/pkg/repository"
	"github.com/Manifoldz/TodoListRESTAPI/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error while reading config file: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("error while connecting to database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	svr := new(server.Server)
	if err := svr.Start(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error while starting http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
