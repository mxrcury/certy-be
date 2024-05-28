package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mxrcury/certy/pkg/logger"
)

type response struct {
	Message string `json:"message"`
}

func sendResponse(c *gin.Context, code int, msg string) {
	logger.Error(msg)
	c.AbortWithStatusJSON(code, response{msg})
}
