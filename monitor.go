package main

import (
	"io/ioutil"
	"net/http"
	"time"
	"unsafe"
)

type ServicesJson struct {
	Services []Service `json:"service"`
	Host     string    `json:"host"`
}

type Service struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	IP     string `json:"IP"`
	Port   string `json:"port"`
	Check  Check  `json:"check"`
	status string
}

type Check struct {
	Type     string `json:"type"`
	Interval string `json:"interval"`
	Timeout  string `json:"timeout"`
}

type Monitor struct {
	microservices []Service
	host          string
}

func NewMonitor() *Monitor {
	microservices := GetMicroservices()
	host := GetHostAddr()
	return &Monitor{microservices: microservices, host: host}
}

func (s *Service) httpHeartbeat() {
	log.Info("service http heartbeat")
	log.Info(s.Check.Interval)
	interval, err := time.ParseDuration(s.Check.Interval)
	if err != nil {
		log.Error(err)
		return
	}

	ticker := time.NewTicker(interval)
	url := "http://" + s.IP + ":" + s.Port + "/healthcheck"
	for {
		select {
		case <-ticker.C:
			log.Info("Send heartbeat request to service: ", s.Id)
			timeout, err := time.ParseDuration(s.Check.Timeout)
			if err != nil {
				log.Error(err)
				return
			}

			tr := &http.Transport{
				ResponseHeaderTimeout: timeout,
			}
			client := &http.Client{Transport: tr, Timeout: timeout}

			resp, err := client.Get(url)
			if err != nil {
				log.Error(err)
				s.status = UNHEALTHY
				continue
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Error(err)
				s.status = UNHEALTHY
				continue
			}
			s.status = HEALTHY
			str := *(*string)(unsafe.Pointer(&body))
			log.Info("Receive from: ", s.Id, ", name: ", s.Name, ", Info: ", str)
		}
	}

}

func (s *Service) tcpHeartbeat() {
	log.Info("service tcp heartbeat")
}

func (m *Monitor) Start() {
	services := m.microservices
	for i, _ := range m.microservices {
		go func(service *Service) {
			if service.Check.Type == "http_heartbeat" {
				service.httpHeartbeat()
			} else if service.Check.Type == "tcp_heartbeat" {
				service.tcpHeartbeat()
			}
		}(&services[i])
	}
}
