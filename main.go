package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"hu-design-project/application/handler"
	_ "hu-design-project/docs"
	"hu-design-project/infrastructure/configuration"
	"hu-design-project/infrastructure/controller"
	"hu-design-project/infrastructure/repository"
	"log"
	"path/filepath"
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
	abs, err := filepath.Abs("./")

	// Printing if there is no error
	if err == nil {
		fmt.Println("Absolute path is:", abs)
	}
	// Init Song Handler
	songHandler := handler.InitializeUserHandler(userRepo)

	// Create Controller
	userController := controller.NewUserController(songHandler)
	e := echo.New()
	userController.Register(e)
	e.Logger.Fatal(e.Start(":8080"))
}
