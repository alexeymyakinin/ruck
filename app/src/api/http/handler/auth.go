package handler

import (
	"errors"
	"github.com/alexeymyakinin/ruck/app/src/core/dep"
	"github.com/alexeymyakinin/ruck/app/src/core/helper"
	"github.com/alexeymyakinin/ruck/app/src/core/service"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func SignIn(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	switch {
	case !c.Request().Form.Has("username"):
		return echo.NewHTTPError(http.StatusUnauthorized, "username not provided")
	case !c.Request().Form.Has("password"):
		return echo.NewHTTPError(http.StatusUnauthorized, "password not provided")
	}

	userServ := dep.GetUserService()
	ctx := c.Request().Context()
	usr, err := userServ.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, helper.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "user not found")
		}
		return echo.ErrInternalServerError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid credentials")
	}

	t, err := service.GenerateJWT(usr.ID, usr.Username)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{"token": t})
}
