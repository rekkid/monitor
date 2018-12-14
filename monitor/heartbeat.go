package monitor

import (
	"monitor/util/config"
	"net/http"
	"time"
)

type HeartBeatService struct {
	service      *config.Service
	countTimeout int
	isActive     bool
}

func NewHeartBeatService(interval string) {
	return HeartBeatService{interval:}
}

func (hb *HeartBeatService) Start() {
	url := "http://" + hb.ip + ":" + hb.port + "/heartbeart"
	tr := &http.Transport{ResponseHeaderTimeout: 5 * time.Second,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		log.Error("Time out receive microservice: %v:%v", hb.ip, hb.port)
		hb.countTimeout++
		if hb.countTimeout > hb.service.Check.Interval {
			log.Error("Heartbeat check failed")
			hb.isActive = false
		}
	}
	hb.countTimeout = 0
	hb.isActive = true
	defer resp.Body.Close()
}
