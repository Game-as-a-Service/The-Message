package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HeartbeatHandler struct {
	Engine *gin.Engine
}

func RegisterHeartbeatHandler(opts *HeartbeatHandler) {
	handler := &HeartbeatHandler{}

	opts.Engine.GET("/api/v1/health", handler.Heartbeat)
}

// Heartbeat godoc
// @Summary Check if the server is alive
// @Description Check if the server is alive
// @Tags heartbeat
// @Accept json
// @Produce json
// @Success 204
// @Router /api/v1/health [GET]
func (g *HeartbeatHandler) Heartbeat(c *gin.Context) {
	c.JSON(http.StatusNoContent, http.NoBody)
}
