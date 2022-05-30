package chat

import (
	"errors"
	"net/http"

	"app/src/core/service/chat"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func GetChat(ctx *gin.Context) {
	id := ctx.GetUint64("id")
	svc := chat.NewChatService()
	res, err := svc.GetChat(ctx.Request.Context(), id)
	if err != nil {
		log.Error().Err(err).Uint64("id", id).Caller().Send()
		switch {
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{
				"error": "unknown error",
			})

		case errors.Is(err, chat.ErrIDNotFound):
			ctx.AbortWithStatusJSON(http.StatusNotFound, &gin.H{
				"error": err.Error(),
			})
		}
		return
	}
	ctx.JSON(http.StatusOK, &res)
}

func GetChatMessages(ctx *gin.Context) {
	id := ctx.GetUint64("id")
	page := ctx.GetUint64("page")
	size := ctx.GetUint64("size")

	svc := chat.NewChatService()
	res, err := svc.GetChatMessages(ctx.Request.Context(), id, page, size)
	if err != nil {
		log.Error().Err(err).Uint64("id", id).Uint64("page", page).Uint64("size", size).Caller().Send()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{
			"error": "unknown error",
		})
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
