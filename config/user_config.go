package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

const userIniName = "user_config.ini"

func LoadUserIni() *ini.File {
	cfg, err := ini.Load(userIniName)
	if err != nil {
		fmt.Printf("Fail to read file: %v\n", err)
		check(err)
	}
	return cfg
}
