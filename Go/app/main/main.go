package main

import (
	"github.com/Gleb-afk/Go/Go/internal/config"
	"github.com/Gleb-afk/Go/Go/internal/handlers/auth"
	"github.com/Gleb-afk/Go/Go/pkg/logging"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initialized")

	logger.Println("config initializing")
	cfg := config.GetConfig()

	logger.Println("router initializing")
	router := httprouter.New()

	logger.Println("create and register handlers")

	userService := service.NewService(cfg.UserService.URL, "/users", logger)
	authHandler := auth.Handler{JWTHelper: jwtHelper, UserService: userService, Logger: logger}
	authHandler.Register(router)

	logger.Println("start application")
}
