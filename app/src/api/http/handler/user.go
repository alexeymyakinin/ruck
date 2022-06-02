package handler

import (
	"github.com/alexeymyakinin/ruck/app/src/api/http/schema"
	"github.com/alexeymyakinin/ruck/app/src/core/dep"
	"github.com/alexeymyakinin/ruck/app/src/core/helper"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func CreateUser(c echo.Context) error {
	var user schema.UserCreateRequest
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "cannot parse body")
	}

	ctx := c.Request().Context()
	svc := dep.GetUserService()
	res, err := svc.CreateUser(ctx, &user)
	if err != nil {
		return helper.GetHTTPError(err)
	}
	return c.JSON(http.StatusCreated, schema.NewUserCreateResponse(res.ID, res.Username))
}

func GetUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			echo.NewBindingError("id", c.ParamValues(), "cannot parse id", err),
		)
	}

	ctx := c.Request().Context()
	svc := dep.GetUserService()
	usr, err := svc.GetUser(ctx, id)
	if err != nil {
		return helper.GetHTTPError(err)
	}

	return c.JSON(http.StatusOK, schema.NewUserSimpleResponse(usr.ID, usr.Username, nil, schema.UserOffline))
}
