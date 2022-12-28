package main

import (
	"dm/dm/client/gmongoclient"
	"dm/dm/client/nacosclient"
	"dm/dm/internal/config"
	"dm/dm/internal/service"
	"dm/dm/setting"
	"dm/pkg/eventCenter"
)

func init() {
	// 解析命令行参数
	_, httpMode, _ := config.ParseFlag()
	// 系统启动setting
	setting.Setup(httpMode)
	gmongoclient.Setup()
	nacosclient.SetUp()
	// nacos监听初始化
	config.InitConfigs()
	config.NacosConfigSetUp()
	// 事件中心生产者初始化
	eventCenter.SetUp(setting.EventKafkaSetting)
}

func main() {
	service.Sync()
}
