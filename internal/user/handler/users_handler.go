package handler

import (
	"net/http"
	"strconv"

	"Back-end/helper"
	"Back-end/internal/user/adapter"
	"Back-end/internal/user/model"

	"github.com/labstack/echo/v4"
)

type EchoUserController struct {
	Service adapter.AdapterUserService
}

// GetUsersController godoc
// @Summary      Get All User
// @Description  Admin can get all users information
// @Tags         User
// @accept       json
// @Produce      json
// @Router       /user [get]
// @Success      200  {object}  model.User
// @Security     JWT
func (ce *EchoUserController) GetUsersController(c echo.Context) error {
	users := ce.Service.GetAllUsersService()

	return c.JSONPretty(http.StatusOK, users, " ")
}

// GetUserController godoc
// @Summary      Get User by Id
// @Description  Admin can get user information by id
// @Tags         User
// @accept       json
// @Produce      json
// @Router       /user/{id} [get]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.User
// @Failure      404  {object}  model.User
// @Security     JWT
func (ce *EchoUserController) GetUserController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	user, err := ce.Service.GetUserByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "no id",
		})
	}

	return c.JSONPretty(http.StatusOK, user, " ")
}

func (ce *EchoUserController) GetUserByTokenController(c echo.Context) error {
	type Input struct {
		Token string `json:"token" form:"token"`
	}
	input := Input{}
	c.Bind(&input)

	result, err := helper.ExtractToken(input.Token)
	if !err {
		return c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"message": "no token",
		})
	}
	return c.JSONPretty(http.StatusOK, result, " ")

	// token := c.Get("token").(*jwt.Token)
	// var result int

	// if token != nil && token.Valid {
	// 	claims := token.Claims.(jwt.MapClaims)
	// 	ID := claims["id"]

	// 	switch ID.(type) {
	// 	case float64:
	// 		result = int(ID.(float64))
	// 	default:
	// 		result = ID.(int)
	// 	}
	// }

	// return c.JSONPretty(http.StatusOK, result, " ")
}

// UpdateUserController godoc
// @Summary      Update User
// @Description  User can update their status or information
// @Tags         User
// @accept       json
// @Produce      json
// @Router       /user/{id} [put]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.User
// @Failure      404  {object}  model.User
// @Security     JWT
func (ce *EchoUserController) UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	user := model.User{}
	c.Bind(&user)

	err := ce.Service.UpdateUserByIDService(intID, user)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "edited",
		"id":      intID,
		"user":    user,
	})
}

// DeleteUsercontroller godoc
// @Summary      Delete User
// @Description  Admin or User can delete their own account
// @Tags         User
// @accept       json
// @Produce      json
// @Router       /user/{id} [delete]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.User
// @Failure      404  {object}  model.User
// @Security     JWT
func (ce *EchoUserController) DeleteUsercontroller(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Service.DeleteUserByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted",
	})
}

// ChangeUsernameController godoc
// @Summary      Update Username
// @Description  User can change username
// @Tags         User
// @accept       json
// @Produce      json
// @Router       /update/username/{id} [put]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.User
// @Failure      404  {object}  model.User
// @Security     JWT
func (ce *EchoUserController) ChangeUsernameController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	user := model.User{}
	c.Bind(&user)

	username := user.Username
	err := ce.Service.UpdateUsernameService(intID, username)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "updated",
	})
}

// ChangePasswordController godoc
// @Summary      Password
// @Description  User can change username
// @Tags         User
// @accept       json
// @Produce      json
// @Router       /update/password/{id} [put]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.User
// @Failure      404  {object}  model.User
// @Security     JWT
func (ce *EchoUserController) ChangePasswordController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	user := model.User{}
	c.Bind(&user)

	password := user.Password
	err := ce.Service.UpdatePasswordService(intID, password)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "updated",
	})
}
