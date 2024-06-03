package controller

import (
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hu-design-project/application/handler/user"
	"hu-design-project/application/model"
	"net/http"
	"net/url"
)

type UserController struct {
	userHandler *user.Handler
}

func NewUserController(userHandler *user.Handler) *UserController {
	return &UserController{
		userHandler: userHandler,
	}
}

func (controller *UserController) Register(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/users", controller.CreateUser)
	e.GET("/users/bulk/:id", controller.GetUserBulk)
	e.POST("/users/activate/:email", controller.ActivateUser)
	e.PATCH("/users/changepassword", controller.ChangePassword)
	e.POST("/users/forgetpassword", controller.ForgetPassword)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags user
// @Accept  json
// @Produce  json
// @Param userCreateModel body model.UserCreateModel true "Create User Model"
// @Success 201 {object} dto.UserDTO "User successfully created"
// @Failure 400 {object} object "Invalid user data"
// @Failure 500 {object} object "Internal Server Error"
// @Router /users [post]
func (controller *UserController) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()
	userCreateModel := new(model.UserCreateModel)
	log.Info("[UserController] creating user")

	if err := c.Bind(userCreateModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user data")
	}

	response, err := controller.userHandler.CreateUser.Handle(ctx, userCreateModel)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, response)
}

// GetUserBulk godoc
// @Summary Get user by ID
// @Description Get details of a user by user ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} mongo_model.User "Successful operation"
// @Failure 400 {string} string "Invalid ID format"
// @Failure 404 {string} string "User not found"
// @Router /users/bulk/{id} [get]
func (controller *UserController) GetUserBulk(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	log.Info("[UserController] getting user")

	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID format")
	}

	response, err := controller.userHandler.GetUser.Handle(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, response)
}

// ActivateUser godoc
// @Summary Activate user
// @Description Activates a user account with the given email
// @Tags user
// @Accept  json
// @Produce  json
// @Param  email path string true "User Email"
// @Success 200 {object} map[string][]string
// @Failure 404 {object} map[string][]string
// @Failure 500 {object} map[string][]string
// @Router /users/activate/{email} [post]
func (controller *UserController) ActivateUser(c echo.Context) error {
	ctx := c.Request().Context()
	encodedEmail := c.Param("email")
	email, err := url.QueryUnescape(encodedEmail)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid email format")
	}
	log.Info("[UserController] Activating user with email:", email)

	err = controller.userHandler.ActivateUser.Handle(ctx, email)
	if err != nil {
		log.Info("[UserController] Error occurred when activating user:", err)
		if err.Error() == "user not found" {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Error occurred when activating user")
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User successfully activated"})
}

// ChangePassword godoc
// @Summary Change user password
// @Description Change the password of a user by verifying the old password and setting a new one
// @Tags user
// @Accept  json
// @Produce  json
// @Param changePassword body model.ChangePasswordModel true "Change Password Information"
// @Success 204 "Password successfully changed"
// @Failure 400 {object} object "Invalid input provided"
// @Failure 500 {object} object "Internal server error"
// @Router /users/changepassword [patch]
func (controller *UserController) ChangePassword(c echo.Context) error {
	var changePassword model.ChangePasswordModel
	if err := c.Bind(&changePassword); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	// Kullanıcının eski şifresini doğrulama ve yeni şifre ile güncelleme işlemi
	if err := controller.userHandler.ChangePassword.Handle(c.Request().Context(), changePassword); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// ForgetPassword godoc
// @Summary Send password reset email
// @Description Sends a password reset email to the user if the email exists in the system
// @Tags user
// @Accept json
// @Produce json
// @Param email query string true "User Email"
// @Success 200 {string} string "Password reset email sent successfully"
// @Failure 400 {string} string "Email is required"
// @Failure 500 {string} string "Internal server error"
// @Router /users/forgetpassword [post]
func (controller *UserController) ForgetPassword(c echo.Context) error {
	email := c.QueryParam("email")
	if email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Email is required")
	}

	// E-posta adresine göre kullanıcıyı bulma ve şifre sıfırlama işlemi
	err := controller.userHandler.ResetPassword.Handle(c.Request().Context(), email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Password reset email sent successfully")
}
