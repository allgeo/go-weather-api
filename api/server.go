package api

import (
	"fmt"
	"log"
	"os"

	weather "blog-apis/api/clients/weatherservice"
	"blog-apis/api/controllers"
	"blog-apis/api/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func getServicer() *controllers.Service {

	weatherBaseURL := os.Getenv("WEATHER_BASE_URL")
	weatherApiKey := os.Getenv("WEATHER_API_KEY")

	if weatherBaseURL == "" || weatherApiKey == "" {
		log.Panic("Either WEATHER_BASE_URL/WEATHER_API_KEY is empty")

	}

	weatherclient, err := weather.NewWeatherServiceClient(weatherBaseURL, weatherApiKey)

	if err != nil {
		log.Panic("Error while create weather client")
	}

	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())

	return controllers.NewService(weatherclient, router)

}

func Run() {
	server := getServicer()
	controllers.NewHttpHandler(server)
	apiPort := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	fmt.Printf("Listening to port %s", apiPort)

	server.Run(apiPort)
}
