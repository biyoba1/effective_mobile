package main

import (
	"github.com/biyoba1/effective_mobile"
	"github.com/biyoba1/effective_mobile/initializer"
	"github.com/biyoba1/effective_mobile/internal/handler"
	"github.com/biyoba1/effective_mobile/internal/repository"
	services2 "github.com/biyoba1/effective_mobile/internal/services"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
	initializer.SyncDatabase()
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	repos := repository.NewRepository(initializer.DB)
	services := services2.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(effective_mobile.Server)
	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Failed to start a server: %s", err.Error())
	}
}
