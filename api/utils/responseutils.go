package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(r *gin.Context, statuscode int, message string) {
	r.JSON(statuscode, gin.H{
		"status": statuscode,
		"error":  message,
	})
}

func OkResponse(r *gin.Context, response interface{}) {
	r.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": response,
	})
}

func OkResponseMessage(r *gin.Context, message string) {
	r.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
