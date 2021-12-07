package module

import (
	"flag"
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"sync"
)

type Config struct {
	InfluxdbToken  string
	InfluxdbUrl    string
	InfluxdbBucket string
	InfluxdbOrg    string
	Lighthouse     []Lighthouse
}
type Lighthouse struct {
	Secretid  string
	Secretkey string
	Regions   []Regions
}
type Regions struct {
	Region string
}

var Configs *Config
var once sync.Once
var Client influxdb2.Client

func init() {
	once.Do(func() {
		conf := flag.String("conf", "./conf/config.yml", "config file path,default ./conf/config.yaml")
		flag.Parse()
		data, _ := ioutil.ReadFile(*conf)
		err := yaml.Unmarshal(data, &Configs)
		if err != nil {
			panic("decode error")
		}
		Client = ConnectInfluxdb(Configs)
	})
}
