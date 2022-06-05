package influxdb

import (
	"github.com/ajski1701/go-meater-meter/config"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func GetInfluxClient() influxdb2.Client {
	token := config.LoadConfigIni().Section("influxdb").Key("token").String()
	url := config.LoadConfigIni().Section("influxdb").Key("url").String()
	return influxdb2.NewClient(url, token)
}
