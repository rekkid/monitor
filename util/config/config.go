package config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Host          string          `json:"host"`
	Microservices []Microservices `json:"microservices"`
}

type Microservices struct {
	Name string `json:"name"`
	IP   string `json:"IP"`
	Port string `json:"prot"`
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

func GetMicroservices() []Microservices {
	return configuration.Microservices
}

func GetHostAddr() string {
	return configuration.Host
}
