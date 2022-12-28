package config

import (
	"dm/dm/internal/model"
	"dm/pkg/myerr"
	"dm/pkg/util"
)

var (
	smallRankCfg []model.SmallRankStruct
)

func RefreshSmallRankCfg(data, md5 string) *myerr.MyErr {
	var smallRankTmpCfg []model.SmallRankStruct
	// map格式的活动配置
	NacosConfigUnmarshal(data, &smallRankTmpCfg)
	smallRankCfg = smallRankTmpCfg
	return myerr.SUCCESS
}

func GetSmallRankSid() (int, int, int) {
	nTime := util.Ntime()
	for _, v := range smallRankCfg {
		if v.StartTime <= nTime && v.EndTime >= nTime {
			return v.Sid, v.StartTime, v.EndTime
		}
	}
	return 0, 0, 0
}
