package main

type ServicesJson struct {
	Services []Service `json:"service"`
	Host     string    `json:"host"`
}

type Service struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	IP    string `json:"IP"`
	Port  string `json:"port"`
	Check Check  `json:"check"`
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
}

func (s *Service) tcpHeartbeat() {
	log.Info("service tcp heartbeat")
}

func (m *Monitor) Start() {
	for _, service := range m.microservices {
		if service.Check.Type == "http_heartbeat" {
			service.httpHeartbeat()
		} else if service.Check.Type == "tcp_heartbeat" {
			service.tcpHeartbeat()
		}
	}
}
