package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHttpHandler(s *Service) {

	handler := &HTTPHandler{s}
	v1 := s.Router.Group("/api")
	{
		// Health Check
		v1.GET("/ping", HealthCheck)
		// Get current weather
		v1.GET("/weather/:city", handler.GetCurrentWeather)

	}
}

func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}
