package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"Back-end/internal/adapter"
	"Back-end/internal/model"
)

type EchoRoleController struct {
	Service adapter.AdapterRoleService
}

// CreateRoleController godoc
// @Summary      Create Role for User
// @Description  Role to identify whether the login is admin or user
// @Tags         Role
// @accept       json
// @Produce      json
// @Router       /role [post]
// @param        data  body      model.Role  true  "required"
// @Success      200   {object}  model.Role
// @Failure      400   {object}  model.Role
// @Failure      500   {object}  model.Role
// @Security     JWT
func (ce *EchoRoleController) CreateRoleController(c echo.Context) error {
	role := model.Role{}
	c.Bind(&role)

	err := ce.Service.CreateRoleService(role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"role":     role,
	})
}

// GetRolesController godoc
// @Summary      Get All Roles Information
// @Description  To get all information about roles
// @Tags         Role
// @accept       json
// @Produce      json
// @Router       /role [get]
// @Success      200  {object}  model.Role
// @Failure      400  {object}  model.Role
// @Security     JWT
func (ce *EchoRoleController) GetRolesController(c echo.Context) error {
	roles := ce.Service.GetAllRolesService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"roles":    roles,
	}, " ")
}

// GetRoleController godoc
// @Summary      Get Role Information by ID
// @Description  Admin can get role information by ID
// @Tags         Role
// @accept       json
// @Produce      json
// @Router       /role/{id} [get]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.Role
// @Failure      400  {object}  model.Role
// @Security     JWT
func (ce *EchoRoleController) GetRoleController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	res, err := ce.Service.GetRoleByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"role":     res,
	})
}

// UpdateRoleController godoc
// @Summary      Update Role by ID
// @Description  Admin can update about role information
// @Tags         Role
// @accept       json
// @Produce      json
// @Router       /role/{id} [put]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.Role
// @Failure      400  {object}  model.Role
// @Failure      500  {object}  model.Role
// @Security     JWT
func (ce *EchoRoleController) UpdateRoleController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	role := model.Role{}
	c.Bind(&role)

	err := ce.Service.UpdateRoleByIDService(intID, role)
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

// DeleteRoleController godoc
// @Summary      Delete Role by ID
// @Description  Admin can delete role information if role information is wrong
// @Tags         Role
// @accept       json
// @Produce      json
// @Router       /role/{id} [delete]
// @param        id   path      int  true  "id"
// @Success      200  {object}  model.Role
// @Failure      400  {object}  model.Role
// @Failure      500  {object}  model.Role
// @Security     JWT
func (ce *EchoRoleController) DeleteRoleController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Service.DeleteRoleByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "deleted",
	})
}
