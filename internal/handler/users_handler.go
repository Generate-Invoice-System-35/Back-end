package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"Back-end/internal/adapter"
	"Back-end/internal/model"
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
// @Failure      400  {object}  model.User
// @Security     JWT
func (ce *EchoUserController) GetUsersController(c echo.Context) error {
	users := ce.Service.GetAllUsersService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    users,
	}, " ")
}

// GetUserController godoc
// @Summary      Get User By Id
// @Description  Admin Can Get User Information By Id
// @Tags         User
// @accept       json
// @Produce      json
// @Router       /user/{id} [get]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.User
// @Failure      400  {object}  model.User
// @Security     JWT
func (ce *EchoUserController) GetUserController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	res, err := ce.Service.GetUserByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    res,
	})
}

// UpdateUserController godoc
// @Summary      Update User
// @Description  User can Update their status or information
// @Tags         User
// @accept       json
// @Produce      json
// @Router       /user/{id} [put]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.User
// @Failure      400  {object}  model.User
// @Failure      500  {object}  model.User
// @Security     JWT
func (ce *EchoUserController) UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	user := model.User{}
	c.Bind(&user)

	err := ce.Service.UpdateUserByIDService(intID, user)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
		"id":       intID,
	})
}

// DeleteUsercontroller godoc
// @Summary      Delete User
// @Description  Admin or User can Delete their own account
// @Tags         User
// @accept       json
// @Produce      json
// @Router       /user/{id} [delete]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.User
// @Failure      400  {object}  model.User
// @Failure      500  {object}  model.User
// @Security     JWT
func (ce *EchoUserController) DeleteUsercontroller(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Service.DeleteUserByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "deleted",
	})
}
