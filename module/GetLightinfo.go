package module

import (
	"github.com/prometheus/common/log"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

func GetLighthoustInfo() {
	var traffic int64
	var instanceid string
	for _, i := range Configs.Lighthouse {
		log.Info("开始获取流量信息")
		credential := common.NewCredential(i.Secretid, i.Secretkey)
		for _, sk := range i.Regions {
			client, err := lighthouse.NewClient(credential, sk.Region, profile.NewClientProfile())
			if err != nil {
				log.Error("与腾讯云建立连接出错", err)
				continue
			}
			request := lighthouse.NewDescribeInstancesTrafficPackagesRequest()
			resp, err := client.DescribeInstancesTrafficPackages(request)
			if err != nil {
				log.Error("获取实例流量出错", err)
				continue
			}
			for _, i := range resp.Response.InstanceTrafficPackageSet {
				for _, m := range i.TrafficPackageSet {
					x := m
					traffic += *x.TrafficUsed + *x.TrafficOverflow
				}
				instanceid = *i.InstanceId
			}
			WritePoint(instanceid, formatFileSize(float64(traffic)))
			traffic = 0
		}
	}
}
func formatFileSize(fileSize float64) (size float64) {
	if fileSize < 1024 {
		return fileSize / float64(1)
	} else if fileSize < (1024 * 1024) {
		return fileSize / float64(1024) * 1000
	} else if fileSize < (1024 * 1024 * 1024) {
		return fileSize / float64(1024*1024) * (1000 * 1000)
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fileSize / float64(1024*1024*1024) * (1000 * 1000 * 1000)
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fileSize / float64(1024*1024*1024*1024) * (1000 * 1000 * 1000 * 1000)
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fileSize / float64(1024*1024*1024*1024*1024) * (1000 * 1000 * 1000 * 1000 * 1000)
	}
}
