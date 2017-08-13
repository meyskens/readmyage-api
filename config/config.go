package config

import (
	"encoding/json"
	"os"
)

// Config contains the info of the configfile
type Config struct {
	Bind      string `json:"bind"`
	Hostname  string `json:"hostname"`
	AutoTLS   bool   `json:"autoTLS"`
	CertCache string `json:"certCache"`
}

// GetConfig reads the config file from disk on ./config.json
func GetConfig() Config {
	returnConfig := Config{}

	data, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	jsonParser := json.NewDecoder(data)
	err = jsonParser.Decode(&returnConfig)
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port != "" {
		returnConfig.Bind = ":" + port
	}
	return returnConfig
}
