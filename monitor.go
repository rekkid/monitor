package main

import (
	"monitor/util/config"
)

//var monitor *Monitor

type Monitor struct {
	microservices []config.Service
	host          string
}

func NewMonitor() *Monitor {
	microservices := config.GetMicroservices()
	host := config.GetHostAddr()
	return &Monitor{microservices: microservices, host: host}
}

func MonitorService(microservice *config.Service, monitorType string) {
	if monitorType == "heartbeat" {
		HeartBeatMonitor()
	}
}

func HeartBeatMonitor() {

}

func (m *Monitor) Start() {

	//for microservice := range microservices {
	//	for monitorType := range monitorTypes {
	//
	//	}
	//}
}
