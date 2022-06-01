package handler

import (
	"app/src/api/http/schema"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"app/src/core/dep"
)

func GetChat(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, schema.NewError(err, "cannot get id"))
	}

	svc := dep.Dep.ChatService()
	res, err := svc.GetChat(c.Request().Context(), id)
	switch {
	case err != nil:
		return c.JSON(http.StatusInternalServerError, schema.NewError(err, ""))
	case res == nil:
		return c.JSON(http.StatusNotFound, schema.NewError(nil, "chat not found"))
	default:
		return fmt.Errorf("")
	}

}
