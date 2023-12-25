package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"hu-design-project/infrastructure/configuration"
	"hu-design-project/infrastructure/repository"
	"log"
	"net/http"
)

func main() {

	config, err := configuration.ReadMongoConfig()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	repo := repository.NewUserRepository()
	err = repo.ConnectToMongo(config.URL)
	if err != nil {
		fmt.Printf("Failed to connect MongoDB with url: %s", config.URL)
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Merhaba, Echo ile API'ye hoş geldiniz!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
