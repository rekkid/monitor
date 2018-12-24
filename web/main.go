package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rekkid/monitor/util/zjlog"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/api", apiHandler)
	return router
}

func hello() {
	log.Info("teat")
}

func main() {
	logfile := "log/log_" + time.Now().Format("2006-01-02") + ".txt"
	var err error
	log, err = util.NewLogger("DEBUG", true, logfile)
	if err != nil {
		panic(err)
	}
	defer log.Sync()
	log.Info("Start web...")
	log.Info("Open http://127.0.0.1:20002")
	handles := RegisterHandlers()
	hello()
	http.ListenAndServe(":20002", handles)
}
