package handler

import (
	"net/http"
	"strconv"

	"app/src/api/http/schema"
	"app/src/core/dep"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CreateUser(c *gin.Context) {
	reqId := requestid.Get(c)
	req, err := schema.GetUserCreateRequest(c)

	svc := dep.Dep.UserService()
	res, err := svc.CreateUser(c.Request.Context(), req)

	if err != nil {
		log.Error().Str("req_id", reqId).Err(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, &res)

}

func GetUserByID(c *gin.Context) {
	reqId := requestid.Get(c)

	id, ok := c.Params.Get("id")
	if !ok {
		log.Warn().Str("req_id", reqId).Msg("cannot get id")
		c.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"error": "id not specified"})
		return
	}

	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Warn().Str("req_id", reqId).Str("id", id).Msg("cannot parse id")
		c.AbortWithStatusJSON(http.StatusBadRequest, &gin.H{"error": "id is not uint64"})
		return
	}

	ctx := c.Request.Context()
	svc := dep.Dep.UserService()
	usr, err := svc.GetUserByID(ctx, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "unknown error"})
		return
	}

	if usr == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, &gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, usr)
}
