package controller

import (
	"github.com/labstack/echo/v4"
	"hu-design-project/application/model/mongo_model"
	"hu-design-project/application/repository"
	"net/http"
	"sync"
)

type DashboardController struct {
	alertRepository     repository.AlertsRepository
	blacklistRepository repository.BlacklistRepository
	deviceRepository    repository.DevicesRepository
	requestRepository   repository.RequestRepository
	whitelistRepository repository.WhitelistRepository
}

func NewDashboardController(alertsRepository repository.AlertsRepository, blacklistRepository repository.BlacklistRepository,
	devicesRepository repository.DevicesRepository,
	whitelistRepository repository.WhitelistRepository, requestRepository repository.RequestRepository) *DashboardController {
	return &DashboardController{
		alertRepository:     alertsRepository,
		blacklistRepository: blacklistRepository,
		whitelistRepository: whitelistRepository,
		requestRepository:   requestRepository,
		deviceRepository:    devicesRepository,
	}
}

func (controller *DashboardController) Register(e *echo.Echo) {
	e.GET("/dashboard", controller.GetDashboard)
}

// GetDashboard godoc
// @Summary Retrieve dashboard data
// @Description Retrieves combined data of alerts, blacklists, devices, requests, and whitelists for the dashboard.
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Success 200 {object} mongo_model.DashboardData "Successful retrieval of dashboard data"
// @Failure 500 {string} string "Internal Server Error"
// @Router /dashboard [get]
func (controller *DashboardController) GetDashboard(c echo.Context) error {
	var wg sync.WaitGroup
	var err error

	dashboard := mongo_model.DashboardData{}

	wg.Add(5)

	go func() {
		defer wg.Done()
		alerts, err := controller.alertRepository.GetAllAlerts(c.Request().Context())
		if err != nil {
			c.Logger().Error(err)
		} else {
			dashboard.Alerts = *alerts
		}
	}()

	go func() {
		defer wg.Done()
		blacklists, err := controller.blacklistRepository.GetAllBlacklists(c.Request().Context())
		if err != nil {
			c.Logger().Error(err)
		} else {
			dashboard.Blacklists = *blacklists
		}
	}()

	go func() {
		defer wg.Done()
		devices, err := controller.deviceRepository.GetAllDevices(c.Request().Context())
		if err != nil {
			c.Logger().Error(err)
		} else {
			dashboard.Devices = *devices
		}
	}()

	go func() {
		defer wg.Done()
		requests, err := controller.requestRepository.GetAllRequests(c.Request().Context())
		if err != nil {
			c.Logger().Error(err)
		} else {
			dashboard.Requests = *requests
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		whitelists, err := controller.whitelistRepository.GetAllWhitelists(c.Request().Context())
		if err != nil {
			c.Logger().Error(err)
		} else {
			dashboard.Whitelists = *whitelists
		}
	}()

	wg.Wait()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to gather dashboard data")
	}

	return c.JSON(http.StatusOK, dashboard)
}
