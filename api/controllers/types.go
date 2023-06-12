package controllers

import (
	weather "blog-apis/api/clients/weatherservice"

	"github.com/gin-gonic/gin"
)

type Service struct {
	WeatherService *weather.WeatherServiceClient
	Router         *gin.Engine
}

type HTTPHandler struct {
	service *Service
}

//NewFeedService creates and returns the feedservice
func NewService(wc *weather.WeatherServiceClient, router *gin.Engine) *Service {
	return &Service{wc, router}
}
