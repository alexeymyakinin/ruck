package helper

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

var ErrNotFound = errors.New("not found")

func HandleRepoErr(err error) error {
	switch err {
	default:
		return err
	case sql.ErrNoRows:
		return ErrNotFound
	}
}

func HandleServiceErr(err error) *echo.HTTPError {
	switch err {
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	case ErrNotFound:
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
}
