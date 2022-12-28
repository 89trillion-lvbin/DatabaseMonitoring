package model

type ModuleBasic struct {
	StartTime    int `json:"startTime"`    // 开始时间
	Continue     int `json:"continue"`     // 持续几天
	Interval     int `json:"interval"`     // 休战几天
	UnitSecond   int `json:"unitSecond"`   // 一天多少s
	DelaySec     int `json:"delaySec"`     // 延时领取奖励
	CfgSidNum    int `json:"cfgSidNum"`    // 加载几个赛季的配置
	GiftContinue int `json:"giftContinue"` // 模式礼包持续几天
}
