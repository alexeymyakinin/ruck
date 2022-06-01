package handler

import (
	"errors"
	"net/http"
	"strconv"

	"app/src/core/dep"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func GetChat(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Params.ByName("id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{
			"error": "cannot parse id",
		})
	}

	svc := dep.Dep.ChatService()
	res, err := svc.GetChat(ctx.Request.Context(), id)

	switch {
	case errors.Is(err, pgx.ErrNoRows):
		ctx.AbortWithStatusJSON(http.StatusNotFound, &gin.H{
			"error": "chat not found",
		})
		return
	case err != nil:
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{
			"error": "unknown error",
		})
		return
	default:
		ctx.JSON(http.StatusOK, &res)
		return
	}
}
