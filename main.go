package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"monitor/monitor"
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
	handles := RegisterHandlers()
	monitor := monitor.NewMonitor()
	monitor.Start()
	http.ListenAndServe(":20001", handles)
}
