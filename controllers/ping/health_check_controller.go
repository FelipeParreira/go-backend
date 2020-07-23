package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthCheck godoc
// @Tags health checks
// @Summary Check API health
// @Description ping API for 200-OK response
// @Produce plain
// @Success 200 {string} string
// @Router /health_check [get]
func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK\n")
}
