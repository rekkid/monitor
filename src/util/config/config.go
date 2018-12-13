package config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	WebAddr string `json:"web_addr"`
}

var configuration *Configuration

func init() {
	file, err := os.Open("./config.json")
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

func GetWebAddr() string {
	return configuration.WebAddr
}
