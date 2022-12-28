package commonmodel

import "time"

type MongoBase struct {
	ConnTimeout     time.Duration
	SocketTimeout   time.Duration
	MaxPoolSize     uint64
	MaxConnIdleTime time.Duration
}

type Nacos struct {
	NacosURI      string
	NacosPort     uint64
	LogDir        string
	CacheDir      string
	LogLevel      string
	ConfigGroupId string
	Cluster       string
	Weight        float64
}

// 奖励详情
type RewardDetail struct {
	Type   int `json:"type"`             // 类型
	ID     int `json:"id"`               // 无后缀具体id， 比如道具id、士兵id(5位),英雄id(5位）,不支持获得指定星级/等级的英雄，请走其他接口
	Count  int `json:"count"`            // 数量
	Limit  int `json:"limit,omitempty"`  // 上限
	Reason int `json:"reason,omitempty"` // 来源
}

type EventKafka struct {
	BootStrapServer        string
	Ack                    int
	BatchSize              int
	AllowAutoTopicCreation bool
	OutTime                time.Duration
	FlushTime              time.Duration
}

type HttpHeader struct {
	Gd  string  `json:"gaid"`    // 谷歌广告id 或 Apple idfa
	Ud  string  `json:"uid"`     // 用户uuid
	Vc  int     `json:"cvc"`     // 客户端版本号
	Sv  float64 `json:"svc"`     // 系统版本号
	Dv  string  `json:"device"`  // 设备名称
	Nw  string  `json:"network"` // 网络类型 wifi/lte/3g/2g/other
	Sc  string  `json:"simcode"` // sim卡 code
	Lg  string  `json:"lang"`    // 设备语言
	Ls  string  `json:"ls"`      // 用户设置的游戏语言
	Pf  string  `json:"pf"`      // 平台
	IP  string  `json:"ip"`      // client ip
	Cty string  `json:"country"` // 设备ip对应的国家
	Ap  int     `json:"appid"`   // appId
}
