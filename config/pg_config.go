package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

const pgIniName = "pg_config.ini"

func LoadPgIni() *ini.File {
	cfg, err := ini.Load(pgIniName)
	if err != nil {
		fmt.Printf("Fail to read file: %v\n", err)
		check(err)
	}
	return cfg
}
