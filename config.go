package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Services []Service `json:"services"`
	Host     string    `json:"host"`
}

var configuration *Configuration

var configjson = ``

func init() {
	file, err := os.Open("config/config.json")
	if err != nil {
		fmt.Println("read file error: %v", err)
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
