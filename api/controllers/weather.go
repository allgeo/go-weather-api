package controllers

import (
	"blog-apis/api/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *HTTPHandler) GetCurrentWeather(r *gin.Context) {
	type response struct {
		CityName           string  `json:"cityName"`
		CurrentTemperature float32 `json:"currentTemperature"`
	}

	city := r.Param("city")
	weatherResponse, err := s.service.WeatherService.GetCurrentWeather(city)
	if err != nil {
		log.Println("Error while getting weather response: ", err)
		utils.ErrorResponse(r, http.StatusBadRequest, "Error while getting weather response")
		return
	}
	resp := response{
		CityName:           weatherResponse.Name,
		CurrentTemperature: weatherResponse.Main.Temp,
	}
	utils.OkResponse(r, resp)
}

//USING HARDCODED CITY

// func (s *HTTPHandler) GetCurrentWeather(r *gin.Context) {
// 	type response struct {
// 		CityName           string  `json:"cityName"`
// 		CurrentTemperature float32 `json:"currentTemperature"`
// 	}

// 	weatherResponse, err := s.service.WeatherService.GetCurrentWeather()
// 	if err != nil {
// 		log.Println("Error while getting weather response: ", err)
// 		utils.ErrorResponse(r, http.StatusBadRequest, "Error while getting weather response")
// 		return
// 	}
// 	resp := response{
// 		CityName:           weatherResponse.Name,
// 		CurrentTemperature: weatherResponse.Main.Temp,
// 	}
// 	utils.OkResponse(r, resp)

// }