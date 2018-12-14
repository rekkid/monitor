package monitor

import "monitor/util/config"

var monitor *Monitor

type Monitor struct {
}

func NewMonitor() *Monitor {
	microservices := config.GetMicroservices()
	monitorTypes := config.GetHostAddr()
	return &Monitor{}
}

func MonitorService(microservice *config.Microservice, monitorType *config.MonitorType) {
	if monitorType == "heartbeat" {
		HeartBeatMonitor()
	}
}

func (m *Monitor) Start() {

	for microservice := range microservices {
		for monitorType := range monitorTypes {

		}
	}
}
