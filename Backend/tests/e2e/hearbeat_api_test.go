package e2e

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHeartbeatEndpoint(t *testing.T) {
	// Initiate a new gin router
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Set up the heartbeat endpoint
	router.GET("/api/v1/heartbeat", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	// Prepare a new HTTP request
	req, err := http.NewRequest("GET", "/api/v1/heartbeat", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}

	// Create a response recorder
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	// Check if the status code is 204
	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d but got %d", http.StatusNoContent, response.Code)
	}

	// Assert that the response body is empty
	assert.Equal(t, http.StatusNoContent, response.Code)
	assert.Empty(t, response.Body.String())
}
