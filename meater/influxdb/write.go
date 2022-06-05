package influxdb

import (
	"context"
	"go-meater-meter/config"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func WriteData(client influxdb2.Client, tags map[string]string, fields map[string]interface{}) {
	org := config.LoadConfigIni().Section("influxdb").Key("org").String()
	bucket := config.LoadConfigIni().Section("influxdb").Key("bucket").String()
	writeAPI := client.WriteAPIBlocking(org, bucket)

	// tags := map[string]string{
	// 	"tagname1": "tagvalue1",
	// }
	// fields := map[string]interface{}{
	// 	"field1": value,
	// }
	point := influxdb2.NewPoint("meater", tags, fields, time.Now())

	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		log.Fatal(err)
	}
}
