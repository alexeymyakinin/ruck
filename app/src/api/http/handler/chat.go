package handler

import (
	"github.com/alexeymyakinin/ruck/app/src/api/http/schema"
	"github.com/alexeymyakinin/ruck/app/src/core/dep"
	"github.com/alexeymyakinin/ruck/app/src/core/helper"
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
		return helper.GetHTTPError(err)
	}

	return c.JSON(http.StatusCreated, schema.NewChatCreateResponse(res.ID, res.Name))
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
		return helper.GetHTTPError(err)
	}

	return c.JSON(http.StatusOK, schema.NewChatSimpleResponse(res.ID, res.Name))
}
