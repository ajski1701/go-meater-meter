package influxdb

import (
	"go-meater-meter/config"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func WriteData(influxdb2.Client) {
	org := config.LoadConfigIni().Section("influxdb").Key("org").String()
	bucket := config.LoadConfigIni().Section("influxdb").Key("bucket").String()
	_ = org
	_ = bucket

}
