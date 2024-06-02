package controller

import (
	"github.com/labstack/echo/v4"
	"hu-design-project/application/handler/device"
	"net/http"
)

type DeviceController struct {
	deviceHandler *device.Handler
}

func NewDeviceController(deviceHandler *device.Handler) *DeviceController {
	return &DeviceController{
		deviceHandler: deviceHandler,
	}
}

func (controller *DeviceController) Register(e *echo.Echo) {
	e.GET("/devices", controller.GetDevices)
}

// GetDevices godoc
// @Summary Retrieve all devices
// @Description Get all devices from the system
// @Tags devices
// @Accept json
// @Produce json
// @Success 200 {array} mongo_model.Device "List of all devices"
// @Failure 500 {string} string "Internal Server Error"
// @Router /devices [get]
func (controller *DeviceController) GetDevices(c echo.Context) error {
	ctx := c.Request().Context()
	devices, err := controller.deviceHandler.GetDevices.Handle(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not retrieve devices: "+err.Error())
	}
	return c.JSON(http.StatusOK, devices)
}
