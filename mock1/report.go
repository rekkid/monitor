package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/healthcheck", func(context *gin.Context) {
		log.Println("模拟超时")
		time.Sleep(time.Millisecond * 900)
		context.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.Run(":30001")
}
