package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-template/pkg/health"
)

var (
	// DefaultHeaderName default header name
	DefaultHeaderName = "API-Health-Check"

	// DefaultResponseCode default response code
	DefaultResponseCode = http.StatusOK

	// DefaultResponseText default response text
	DefaultResponseText = "ok"

	// DefaultConfig default config
	DefaultConfig = Config{
		HeaderName:   DefaultHeaderName,
		ResponseCode: DefaultResponseCode,
		ResponseText: DefaultResponseText}
)

// Config holds the configuration values
type Config struct {
	HeaderName   string
	ResponseCode int
	ResponseText string
}

// Healthcheck godoc
// @Summary Get server status
// @Produce json
// @Success 200
// @Router /api/healthcheck [get]
func Health(c *gin.Context) {
	check := health.Check{
		AppName: "go-template",
		Version: "1.0",
	}
	c.JSON(http.StatusOK, check)
}
