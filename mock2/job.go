package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/healthcheck", func(context *gin.Context) {
		//context.JSON(http.StatusOK, gin.H{"status": "ok"})
		context.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.Run(":30002")
}
