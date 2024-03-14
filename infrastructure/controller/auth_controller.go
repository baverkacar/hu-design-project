package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"hu-design-project/application/handler/auth"
	"hu-design-project/application/model"
	"net/http"
)

type AuthController struct {
	authHandler *auth.Handler
	validate    *validator.Validate
}

func NewAuthController(authHandler *auth.Handler) *AuthController {
	return &AuthController{
		authHandler: authHandler,
	}
}

func (controller *AuthController) Register(e *echo.Echo) {
	e.POST("/auth/login", controller.Login)
}

// Login godoc
// @Summary Login a user
// @Description Log in with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param loginModel body model.UserLoginModel true "Login Credentials"
// @Success 200 {object} mongo_model.User "Successfully logged in"
// @Failure 400 {string} string "Invalid request format or Validation error"
// @Failure 401 {string} string "Invalid email or password"
// @Router /auth/login [post]
func (controller *AuthController) Login(c echo.Context) error {
	var loginModel model.UserLoginModel
	if err := c.Bind(&loginModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}

	/*if err := controller.validate.Struct(loginModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Validation error: "+err.Error())
	}*/

	log.Info("[AuthController] Login user with email: ", loginModel.Email)
	user, err := controller.authHandler.UserLogin.Handle(c.Request().Context(), loginModel)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}

	return c.JSON(http.StatusOK, user)

}
