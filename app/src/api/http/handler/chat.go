package handler

import (
	"app/src/api/http/schema"
	"app/src/core/dep"
	"app/src/core/helper"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func CreateChat(c echo.Context) error {
	var chat schema.ChatCreateRequest
	err := c.Bind(&chat)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "cannot parse body")
	}

	ctx := c.Request().Context()
	svc := dep.GetChatService()
	res, err := svc.CreateChat(ctx, &chat)
	if err != nil {
		return helper.HandleServiceErr(err)
	}

	return c.JSON(http.StatusCreated, res)
}

func GetChat(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			echo.NewBindingError("id", c.ParamValues(), "cannot parse id", err),
		)
	}

	ctx := c.Request().Context()
	svc := dep.GetChatService()
	res, err := svc.GetChat(ctx, id)
	if err != nil {
		return helper.HandleServiceErr(err)
	}

	return c.JSON(http.StatusOK, res)
}
