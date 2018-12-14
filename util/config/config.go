package config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Services []Service `json:"services"`
	Host     string    `json:"host"`
}

type ServicesJson struct {
	Services []Service `json:"service"`
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

var configuration *Configuration

func init() {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Printf("read file error: %v", err)
		return
	}

	decoder := json.NewDecoder(file)
	configuration = &Configuration{}
	err = decoder.Decode(configuration)
	if err != nil {
		panic(err)
	}
}

func GetMicroservices() []Service {
	return configuration.Services
}

func GetHostAddr() string {
	return configuration.Host
}
