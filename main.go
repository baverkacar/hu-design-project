package main

import (
	"github.com/labstack/echo/v4"
	"hu-design-project/application/handler/auth"
	"hu-design-project/application/handler/user"
	_ "hu-design-project/docs"
	"hu-design-project/infrastructure/configuration"
	"hu-design-project/infrastructure/controller"
	"hu-design-project/infrastructure/repository"
	"log"
)

func main() {

	config, err := configuration.ReadMongoConfig()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	userRepo, err := repository.NewUserRepository(config.URL, config.Database, config.Collection)
	if err != nil {
		log.Fatalf("Failed to connect MongoDB with url: %s, error: %v", config.URL, err)
	}

	// Init Handlers
	userHandler := user.InitializeUserHandler(userRepo)
	authHandler := auth.InitializeAuthHandler(userRepo)

	// Create Controllers
	userController := controller.NewUserController(userHandler)
	authController := controller.NewAuthController(authHandler)

	e := echo.New()
	userController.Register(e)
	authController.Register(e)
	e.Logger.Fatal(e.Start(":8080"))
}
