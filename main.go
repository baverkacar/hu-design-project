package main

import (
	"github.com/labstack/echo/v4"
	"hu-design-project/application/handler/alert"
	"hu-design-project/application/handler/auth"
	"hu-design-project/application/handler/device"
	"hu-design-project/application/handler/request"
	"hu-design-project/application/handler/user"
	_ "hu-design-project/docs"
	"hu-design-project/infrastructure/configuration"
	"hu-design-project/infrastructure/controller"
	"hu-design-project/infrastructure/repository"
	"log"
)

func main() {

	config := configuration.NewMongoConfig()

	userRepo, err := repository.NewUserRepository(config.URL, config.Database, config.UserCollection)
	if err != nil {
		log.Fatalf("Failed to connect MongoDB with url: %s, error: %v", config.URL, err)
	}

	alertsRepo, err := repository.NewAlertsRepository(config.URL, config.Database, config.AlertsCollection)
	if err != nil {
		log.Fatalf("Failed to create AlertsRepository: %v", err)
	}

	devicesRepo, err := repository.NewDevicesRepository(config.URL, config.Database, config.DevicesCollection)
	if err != nil {
		log.Fatalf("Failed to create DevicesRepository: %v", err)
	}

	requestRepo, err := repository.NewRequestRepository(config.URL, config.Database, config.RequestsCollection)
	if err != nil {
		log.Fatalf("Failed to create DevicesRepository: %v", err)
	}

	whitelistRepo, err := repository.NewWhitelistRepository(config.URL, config.Database, config.WhitelistCollection)
	if err != nil {
		log.Fatalf("Failed to create DevicesRepository: %v", err)
	}

	blacklistRepo, err := repository.NewBlacklistRepository(config.URL, config.Database, config.BlacklistCollection)
	if err != nil {
		log.Fatalf("Failed to create DevicesRepository: %v", err)
	}

	// Init Handlers
	userHandler := user.InitializeUserHandler(userRepo)
	authHandler := auth.InitializeAuthHandler(userRepo)
	alertHandler := alert.InitializeAlertHandler(alertsRepo, whitelistRepo, blacklistRepo)
	deviceHandler := device.InitializeDeviceHandler(devicesRepo)
	requestHandler := request.InitializeRequestHandler(requestRepo)

	// Create Controllers
	userController := controller.NewUserController(userHandler)
	authController := controller.NewAuthController(authHandler)
	alertController := controller.NewAlertController(alertHandler, whitelistRepo, blacklistRepo)
	deviceController := controller.NewDeviceController(deviceHandler)
	requestController := controller.NewRequestController(requestHandler)
	dashboardController := controller.NewDashboardController(alertsRepo, blacklistRepo, devicesRepo, whitelistRepo, requestRepo)

	e := echo.New()
	userController.Register(e)
	authController.Register(e)
	alertController.Register(e)
	deviceController.Register(e)
	requestController.Register(e)
	dashboardController.Register(e)
	e.Logger.Fatal(e.Start(":8080"))
}
