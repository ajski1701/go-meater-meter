package main

import (
	"fmt"
	"go-meater-meter/config"
	"go-meater-meter/meater/authentication"
	"go-meater-meter/meater/devices"
	"go-meater-meter/meater/influxdb"
)

func main() {
	authenticationDebug := config.LoadConfigIni().Section("debug").Key("disable_authentication").String()
	sessionToken := ""
	if authenticationDebug == "true" {
		sessionToken = authentication.GetAuth(config.LoadConfigIni())
	} else {
		sessionToken = ""
	}
	devices := devices.GetDevices(sessionToken)
	influxdbClient := influxdb.GetInfluxClient()
	_ = influxdbClient
	for _, device := range devices {
		fmt.Println(device)
	}
}
