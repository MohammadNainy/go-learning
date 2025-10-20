package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events",getHandler)

	server.Run(":8080")
}

func getHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message":"Hello GO!","Status":200})
}