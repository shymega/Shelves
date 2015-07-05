package config

// At the moment, I'm using Viper for configuration.
// This file is here for archival purposes.

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Version string `json:"version"`
	Db      struct {
		Mongodb struct {
			Hostname string `json:"hostname"`
			Port     int    `json:"port"`
		} `json:"mongodb"`
	} `json:"db"`
	Authentication struct {
		Authenticator string `json:"authenticator"`
		File          struct {
			Authdblocation string `json:"authDBlocation"`
		} `json:"file"`
	} `json:"authentication"`
}

var Conf Configuration

func LoadConfig(configLocation string) {
	File, err := os.Open(configLocation)
	if err != nil {
		fmt.Printf("Error occurred opening the configuration file.\nError message: %v\n", err)
	}

	err = json.NewDecoder(File).Decode(&Conf)
	if err != nil {
		fmt.Printf("Error occurred decoding the configuration file into the Configuration struct.\nError message: %v\n", err)
	}
}
