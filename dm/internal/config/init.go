package config

import (
	"github.com/spf13/pflag"
)

// 配置文件路径
const (
	DefaultPort       = 0
	DefaultMode       = ""
	DefaultDaySeconds = 86400
	DefaultGroupId    = ""
)

// 启动参数
var (
	port       = pflag.IntP("port", "", DefaultPort, "")
	mode       = pflag.StringP("mode", "", DefaultMode, "")
	testConfig = pflag.StringP("config", "", "", "iw test config file path")
)

func ParseFlag() (int, string, string) {
	pflag.Parse()

	// mode 只支持3种，参数检验，防止非法mode
	// 添加2种特殊测试配置，后续转换成合法mode
	modes := map[string]struct{}{
		"test5min": {}, // 5分钟测试
		"testFull": {}, // 全量测试
		"debug":    {},
		"test":     {},
		"release":  {},
	}
	if _, ok := modes[*mode]; !ok {
		*mode = ""
	}
	return *port, *mode, *testConfig
}
