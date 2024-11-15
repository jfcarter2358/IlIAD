package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var healthy = false
var ready = false

func Health(ctx *gin.Context) {
	if healthy {
		ctx.JSON(http.StatusOK, gin.H{"status": "healthy"})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{"status": "unhealthy"})
}

func Readiness(ctx *gin.Context) {
	if healthy {
		ctx.JSON(http.StatusOK, gin.H{"status": "ready"})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{"status": "not ready"})
}
