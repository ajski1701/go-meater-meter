package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

const configFileName = "config.ini"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadConfigIni() *ini.File {
	cfg, err := ini.Load(configFileName)
	if err != nil {
		fmt.Printf("Fail to read file: %v\n", err)
		check(err)
	}
	return cfg
}
