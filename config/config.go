package config

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadConfigIni() *ini.File {
	configFile := os.Getenv("CONFIG_FILE")
	if len(configFile) == 0 {
		configFile = "config.ini"
	}
	cfg, err := ini.Load(configFile)
	if err != nil {
		fmt.Printf("Fail to read file: %v\n", err)
		check(err)
	}
	return cfg
}
