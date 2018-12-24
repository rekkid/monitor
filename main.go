package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rekkid/monitor/util/zjlog"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

var (
	log *zjlog.Log
	g   errgroup.Group
)

func ShowServiceStatus(m *Monitor, c *gin.Context) {
	result := make([]gin.H, len(m.microservices))
	for i, service := range m.microservices {
		result[i] = gin.H{
			"host":   m.host,
			"name":   service.Name,
			"status": service.status,
		}
	}
	log.Info("testfakfjsakfjksajfksaj")
	c.JSON(http.StatusOK, result)
}

func c(m *Monitor) func(c *gin.Context) {
	f := func(c *gin.Context) {
		ShowServiceStatus(m, c)
	}
	return f
}

func Mon(m *Monitor) http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/monitor", c(m))
	return e
}

func main() {
	logfile := "log/log_" + time.Now().Format("2006-01-02") + ".txt"
	var err error
	log, err = zjlog.NewLogger("DEBUG", true, logfile)
	if err != nil {
		panic(err)
	}
	defer log.Sync()

	monitor := NewMonitor()
	monitor.Start()

	server1 := &http.Server{
		Addr:         ":20001",
		Handler:      Mon(monitor),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server2 := &http.Server{
		Addr:         ":20002",
		Handler:      RunCmd(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server1.ListenAndServe()
	})

	g.Go(func() error {
		return server2.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}
