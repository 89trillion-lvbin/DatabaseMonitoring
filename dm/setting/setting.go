package setting

import (
	"fmt"
	"log"
	"time"

	"dm/pkg/commonmodel"

	"github.com/go-ini/ini"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
)

type App struct {
	RunMode string
}

var AppSetting = &App{}

type Mongo struct {
	ApplyURI string
}

var (
	NacosClientConfig  constant.ClientConfig
	NacosServerConfigs []constant.ServerConfig
)

var (
	MongoBaseSetting  = &commonmodel.MongoBase{}
	MongoCoreSetting  = &Mongo{}
	NacosSetting      = &commonmodel.Nacos{}
	EventKafkaSetting = &commonmodel.EventKafka{}
)

// Setup initialize the configuration instance
// nolint:funlen
func Setup(mode string) {
	var err error

	appSettingPath := "dm/conf/sys/app.ini"

	var appCfg *ini.File
	appCfg, err = ini.Load(appSettingPath)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse %s : %v", appSettingPath, err)
	}

	mapTo(appCfg, "app", AppSetting)
	if len(mode) > 0 {
		AppSetting.RunMode = mode
	}

	serverSettingPath := fmt.Sprintf("dm/conf/sys/%s.ini", AppSetting.RunMode)
	if AppSetting.RunMode == "test5min" || AppSetting.RunMode == "testFull" { // 特殊环境配置转化
		AppSetting.RunMode = "test"
	}

	var serverCfg *ini.File
	serverCfg, err = ini.Load(serverSettingPath)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse %s: %v", serverSettingPath, err)
	}

	mapTo(serverCfg, "mongo-base", MongoBaseSetting)
	mapTo(serverCfg, "mongo-core", MongoCoreSetting)
	mapTo(serverCfg, "nacos", NacosSetting)
	mapTo(serverCfg, "event-kafka", EventKafkaSetting)

	MongoBaseSetting.ConnTimeout *= time.Second
	MongoBaseSetting.SocketTimeout *= time.Second
	MongoBaseSetting.MaxConnIdleTime *= time.Second

	NacosClientConfig = constant.ClientConfig{
		NamespaceId:         "", // namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              NacosSetting.LogDir,
		CacheDir:            NacosSetting.CacheDir,
		LogLevel:            NacosSetting.LogLevel,
	}

	NacosServerConfigs = []constant.ServerConfig{{
		IpAddr: NacosSetting.NacosURI,
		Port:   NacosSetting.NacosPort,
	}}
}

// mapTo map section
func mapTo(cfg *ini.File, section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
