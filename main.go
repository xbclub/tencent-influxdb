package main

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/robfig/cron/v3"
	"tencent-influxdb/module"
)

var Parser cron.Parser

func main() {
	var wait = make(chan bool)
	Parser = cron.NewParser(
		cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow,
	)

	c := cron.New(cron.WithParser(Parser))

	//添加定时任务 schedule 为cron表达式
	_, err := c.AddFunc("0 */1 * * * *", module.GetLighthoustInfo)
	if err != nil {
		logs.Error(err)
	}
	c.Start()
	<-wait
}
