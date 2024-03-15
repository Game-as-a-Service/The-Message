package e2e

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHeartbeatEndpoint(t *testing.T) {
	// Initiate a new gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Set up the heartbeat endpoint
	router.GET("/api/v1/health", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	// Prepare a new HTTP request
	req, _ := http.NewRequest("GET", "/api/v1/health", nil)

	// Create a response recorder
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	// Assert that the response body is empty
	assert.Equal(t, http.StatusNoContent, res.Code)
	assert.Empty(t, res.Body.String())
}
