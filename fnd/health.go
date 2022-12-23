package fnd

import (
	"github.com/gin-gonic/gin"
)

func HealthHandler(ctx *gin.Context) {
	ctx.JSON(200, map[string]string{"status": "healthy"})
}
