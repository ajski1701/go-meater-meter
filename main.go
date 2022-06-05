package main

import (
	"fmt"
	"time"

	"github.com/ajski1701/go-meater-meter/devices"
	"github.com/ajski1701/go-meater-meter/influxdb"
)

func main() {
	fmt.Println(time.Now(), "Starting go-meater-meter application")
	fmt.Println(time.Now(), "Authenticating to Meater Cloud API")
	sessionToken := getAuthenticationToken()

	for {
		fmt.Println(time.Now(), "Querying Meater Cloud Device API")
		devices := devices.GetDevices(sessionToken)
		influxdbClient := influxdb.GetInfluxClient()
		pollRate := getPollRate()

		submitInfluxData(devices, sessionToken, influxdbClient)
		fmt.Println(time.Now(), "Successfully wrote influxdb data for all devices")
		time.Sleep(time.Duration(pollRate) * time.Second)
	}
}
