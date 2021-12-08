package module

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/siddontang/go/log"
	"time"
)

func ConnectInfluxdb(configs *Config) influxdb2.Client {
	return influxdb2.NewClient(configs.InfluxdbUrl, configs.InfluxdbToken)
}
func WritePoint(name string, traffic float64) {
	defer func() {
		if err := recover(); err != nil {
			time.Sleep(3 * time.Second)
			ConnectInfluxdb(Configs)
			WritePoint(name, traffic)
		}
	}()
	configs := *Configs
	writeAPI := Client.WriteAPIBlocking(configs.InfluxdbOrg, configs.InfluxdbBucket)
	p := influxdb2.NewPoint("Bandwidth", map[string]string{"Name": name}, map[string]interface{}{"traffic": traffic}, time.Now())
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		log.Error(err)
		return
	}
}
