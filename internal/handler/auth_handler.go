package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"Back-end/internal/adapter"
	"Back-end/internal/model"
)

type EchoAuthController struct {
	Service adapter.AdapterAuthService
}

// RegisterController godoc
// @Summary      Register User
// @Description  People can Register as a User
// @Tags         User
// @accept       json
// @Produce      json
// @Router       /register [post]
// @param        data  body      model.User  true  "required"
// @Success      200   {object}  model.User
// @Failure      400   {object}  model.User
// @Failure      500   {object}  model.User
func (ce *EchoAuthController) RegisterController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	err := ce.Service.RegisterService(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"users":    user,
	})
}

// LoginController godoc
// @Summary      Login User
// @Description  People can Login as a User
// @Tags         User
// @accept       json
// @Produce      json
// @Router       /login [post]
// @Param        data  body      model.User  true  "required"
// @Success      200   {object}  model.User
// @Success      400   {object}  model.User
// @Failure      401   {object}  model.User
// @Failure      500   {object}  model.User
func (ce *EchoAuthController) LoginController(c echo.Context) error {
	userLogin := model.User{}
	c.Bind(&userLogin)

	token, statusCode := ce.Service.LoginService(userLogin.Username, userLogin.Password)
	switch statusCode {
	case http.StatusUnauthorized:
		return c.JSONPretty(http.StatusUnauthorized, map[string]interface{}{
			"messages": "username atau password salah",
		}, " ")
	case http.StatusInternalServerError:
		return c.JSONPretty(http.StatusInternalServerError, map[string]interface{}{
			"messages": "internal server error",
		}, " ")
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"token":    token,
	}, " ")
}
