package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"app/src/api/http/schema"
	"app/src/core/dep"
	"github.com/gin-gonic/gin"
)

func CreateUser(c echo.Context) error {
	svc := dep.Dep.UserService()

	req, err := schema.GetUserCreateRequest(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, schema.NewError(err, "cannot parse body"))
	}

	res, err := svc.CreateUser(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, schema.NewError(err, ""))
	}
	return c.JSON(http.StatusCreated, &res)

}

func GetUserByID(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &gin.H{"error": "id is not uint64"})
	}

	svc := dep.Dep.UserService()
	usr, err := svc.GetUserByID(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &gin.H{"error": "unknown error"})
	}

	if usr == nil {
		return c.JSON(http.StatusNotFound, &gin.H{"error": "user not found"})
	}

	return c.JSON(http.StatusOK, usr)
}
