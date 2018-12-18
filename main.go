package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"monitor/util/zjlog"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Welcome to microservice monitor!\n")
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", Index)
	return router
}

var log *zjlog.Log

func main() {
	logfile := "log/log_" + time.Now().Format("2006-01-02") + ".txt"
	var err error
	log, err = zjlog.NewLogger("DEBUG", true, logfile)
	if err != nil {
		panic(err)
	}
	defer log.Sync()
	log.Info("Start main...")
	log.Info("Open http://127.0.0.1:20001")
	_ = RegisterHandlers()
	monitor := NewMonitor()
	monitor.Start()
	//http.ListenAndServe(":20001", handles)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
