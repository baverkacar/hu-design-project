package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"hu-design-project/application/handler/alert"
	"hu-design-project/application/model/mongo_model"
	"net/http"
)

type AlertController struct {
	alertHandler *alert.Handler
}

func NewAlertController(alertHandler *alert.Handler) *AlertController {
	return &AlertController{
		alertHandler: alertHandler,
	}
}

func (controller *AlertController) Register(e *echo.Echo) {
	e.GET("/alerts", controller.GetAlerts)
	e.POST("/alerts/:id", controller.AddToList)
	e.GET("/whitelist/:id", controller.GetWhitelistEntry)
	e.GET("/blacklist/:id", controller.GetBlacklistEntry)
	e.DELETE("/whitelist/:id", controller.DeleteWhitelistEntry)
	e.DELETE("/blacklist/:id", controller.DeleteBlacklistEntry)
}

// GetAlerts godoc
// @Summary Retrieve all alerts
// @Description Get all alerts from the system
// @Tags alerts
// @Accept json
// @Produce json
// @Success 200 {array} mongo_model.Alert "List of all alerts"
// @Failure 500 {string} string "Internal Server Error"
// @Router /alerts [get]
func (controller *AlertController) GetAlerts(c echo.Context) error {
	ctx := c.Request().Context()
	alerts, err := controller.alertHandler.GetAlerts.Handle(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not retrieve alerts: "+err.Error())
	}
	return c.JSON(http.StatusOK, alerts)
}

// AddToList godoc
// @Summary Add or remove an alert's IP to/from whitelist or blacklist
// @Description Adds the IP address of an alert with the given ID to the whitelist or blacklist based on the 'add' query parameter
// @Tags alerts
// @Accept  json
// @Produce  json
// @Param id path string true "Alert ID"
// @Param add query string true "Action to perform ('whitelist' or 'blacklist')"
// @Success 200 {object} mongo_model.List "IP successfully added to list"
// @Failure 400 "Invalid alert ID or operation"
// @Failure 404 "Alert not found"
// @Failure 500 "Internal Server Error"
// @Router /alerts/{id} [post]
func (controller *AlertController) AddToList(c echo.Context) error {
	alertID := c.Param("id")
	action := c.QueryParam("add")
	if action != "whitelist" && action != "blacklist" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid operation. Supported operations are 'whitelist' and 'blacklist'.")
	}

	ctx := c.Request().Context()
	var response *mongo_model.List // Bu model gelecekte whitelist ve blacklist için genişletilebilir bir model olmalı.
	var err error

	if action == "whitelist" {
		response, err = controller.alertHandler.AddWhitelist.Handle(ctx, alertID)
	} else if action == "blacklist" {
		response, err = controller.alertHandler.AddBlacklist.Handle(ctx, alertID)
	}

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return echo.NewHTTPError(http.StatusNotFound, "Alert not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}

// GetWhitelistEntry godoc
// @Summary Retrieve a whitelist entry
// @Description Get a whitelist entry by ID
// @Tags whitelist
// @Accept json
// @Produce json
// @Param id path string true "Whitelist Entry ID"
// @Success 200 {object} mongo_model.List "Successfully retrieved whitelist entry"
// @Failure 404 "Whitelist entry not found"
// @Failure 500 "Internal Server Error"
// @Router /whitelist/{id} [get]
func (controller *AlertController) GetWhitelistEntry(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	entry, err := controller.alertHandler.GetWhitelist.Handle(ctx, id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return echo.NewHTTPError(http.StatusNotFound, "Whitelist entry not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, entry)
}

// GetBlacklistEntry godoc
// @Summary Retrieve a blacklist entry
// @Description Get a blacklist entry by ID
// @Tags blacklist
// @Accept json
// @Produce json
// @Param id path string true "Blacklist Entry ID"
// @Success 200 {object} mongo_model.List "Successfully retrieved blacklist entry"
// @Failure 404 "Blacklist entry not found"
// @Failure 500 "Internal Server Error"
// @Router /blacklist/{id} [get]
func (controller *AlertController) GetBlacklistEntry(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	entry, err := controller.alertHandler.GetBlacklist.Handle(ctx, id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return echo.NewHTTPError(http.StatusNotFound, "Blacklist entry not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, entry)
}

// DeleteWhitelistEntry godoc
// @Summary Remove a whitelist entry
// @Description Delete a whitelist entry by ID
// @Tags whitelist
// @Accept  json
// @Produce  json
// @Param id path string true "Whitelist Entry ID"
// @Success 200 {string} string "Successfully deleted whitelist entry"
// @Failure 404 "Whitelist entry not found"
// @Failure 500 "Internal Server Error"
// @Router /whitelist/{id} [delete]
func (controller *AlertController) DeleteWhitelistEntry(c echo.Context) error {
	id := c.Param("id")

	ctx := c.Request().Context()
	err := controller.alertHandler.DeleteWhitelist.Handle(ctx, id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Blacklist entry not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "Successfully deleted blacklist entry")
}

// DeleteBlacklistEntry godoc
// @Summary Remove a blacklist entry
// @Description Delete a blacklist entry by ID
// @Tags blacklist
// @Accept  json
// @Produce  json
// @Param id path string true "Blacklist Entry ID"
// @Success 200 {string} string "Successfully deleted blacklist entry"
// @Failure 404 "Blacklist entry not found"
// @Failure 500 "Internal Server Error"
// @Router /blacklist/{id} [delete]
func (controller *AlertController) DeleteBlacklistEntry(c echo.Context) error {
	id := c.Param("id")

	ctx := c.Request().Context()
	err := controller.alertHandler.DeleteBlacklist.Handle(ctx, id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Blacklist entry not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "Successfully deleted blacklist entry")
}
