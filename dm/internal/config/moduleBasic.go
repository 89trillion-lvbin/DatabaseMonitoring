package config

import (
	"dm/dm/internal/model"
	"dm/pkg/myerr"
)

var allModuleBasicCfg = map[string]model.ModuleBasic{}

// 各个模块的开始结束时间
func RefreshModuleBasicData(data, md5 string) *myerr.MyErr {
	tmp := map[string]model.ModuleBasic{}
	NacosConfigUnmarshal(data, &tmp)
	allModuleBasicCfg = tmp
	return myerr.SUCCESS
}

// 获取指定模块的基础配置
func GetModuleBasisMustData(moduleType string) (int, int, int, int) {
	cfg := allModuleBasicCfg[moduleType]
	return cfg.StartTime, cfg.Continue, cfg.Interval, cfg.UnitSecond
}

// 活动第一次开启时间
func GetModuleBasicStartTime(moduleType string) int {
	return allModuleBasicCfg[moduleType].StartTime
}

// 活动持续天数
func GetModuleBasicContinueDay(moduleType string) int {
	return allModuleBasicCfg[moduleType].Continue
}

// 下一次开启活动间隔天数
func GetModuleBasicIntervalDay(moduleType string) int {
	return allModuleBasicCfg[moduleType].Interval
}

// 每日时长
func GetModuleBasicUnitSecond(moduleType string) int {
	return allModuleBasicCfg[moduleType].UnitSecond
}

// 拉取赛季数量
func GetModuleBasicCfgSidNum(moduleType string) int {
	return allModuleBasicCfg[moduleType].CfgSidNum
}

// 活动持续时间
func GetModuleBasicGiftContinueDay(moduleType string) int {
	return allModuleBasicCfg[moduleType].GiftContinue
}

//  领取延时
func GetModuleBasicDelaySec(moduleType string) int {
	return allModuleBasicCfg[moduleType].DelaySec
}
