package main

import (
	"go-meater-meter/config"
	"go-meater-meter/meater/authentication"
	"go-meater-meter/meater/devices"
	"go-meater-meter/meater/influxdb"
	models "go-meater-meter/meater/models"
	"strconv"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func getAuthenticationToken() string {
	authenticationDebug := config.LoadConfigIni().Section("debug").Key("disable_authentication").String()
	sessionToken := ""
	if authenticationDebug == "true" {
		sessionToken = ""
	} else {
		sessionToken = authentication.GetAuth(config.LoadConfigIni())
	}
	return sessionToken
}

func getPollRate() int {
	pollRateStr := config.LoadConfigIni().Section("app-config").Key("poll_rate").String()
	pollRateInt, err := strconv.Atoi(pollRateStr)
	if err != nil {
		panic(err)
	}
	return pollRateInt
}

func submitInfluxData(devices []models.Devices, sessionToken string, client influxdb2.Client) {
	for _, device := range devices {
		tags := map[string]string{
			"device_id": device.Id,
			"cook_id":   device.Cook.Id,
			"cook_name": device.Cook.Name,
		}
		fields := map[string]interface{}{
			"internal_temperature":    device.Temperature.Internal,
			"ambient_temperature":     device.Temperature.Ambient,
			"cook_target_temperature": device.Cook.Temperature.Target,
			"cook_peak_temperature":   device.Cook.Temperature.Peak,
			"cook_elapsed_time":       device.Cook.Time.Elapsed,
			"cook_remaining_time":     device.Cook.Time.Remaining,
		}
		influxdb.WriteData(client, tags, fields)
	}
}

func main() {
	sessionToken := getAuthenticationToken()

	for {
		devices := devices.GetDevices(sessionToken)
		influxdbClient := influxdb.GetInfluxClient()
		pollRate := getPollRate()

		submitInfluxData(devices, sessionToken, influxdbClient)
		time.Sleep(time.Duration(pollRate) * time.Second)
	}
}
