package controller

import (
	"github.com/labstack/echo/v4"
	"hu-design-project/application/handler/request" // Uygun paket yoluyla güncelleyin
	"net/http"
)

type RequestController struct {
	requestHandler *request.Handler
}

func NewRequestController(requestHandler *request.Handler) *RequestController {
	return &RequestController{
		requestHandler: requestHandler,
	}
}

func (controller *RequestController) Register(e *echo.Echo) {
	e.GET("/requests", controller.GetRequests)
}

// GetRequests godoc
// @Summary Retrieve all requests
// @Description Get all requests from the system
// @Tags requests
// @Accept json
// @Produce json
// @Success 200 {array} mongo_model.Request "List of all requests"
// @Failure 500 {string} string "Internal Server Error"
// @Router /requests [get]
func (controller *RequestController) GetRequests(c echo.Context) error {
	ctx := c.Request().Context()
	requests, err := controller.requestHandler.GetRequests.Handle(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not retrieve requests: "+err.Error())
	}
	return c.JSON(http.StatusOK, requests)
}
