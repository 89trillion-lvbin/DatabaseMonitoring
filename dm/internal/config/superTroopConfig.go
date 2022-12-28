package config

import (
	"dm/dm/internal/model"
	"dm/pkg/myerr"
	"dm/pkg/util"
)

var (
	superTroopSeasons []model.SuperTroopSeasonCfg
)

func RefreshSuperTroopRewards(data, md5 string) *myerr.MyErr {
	tmpSeasonsCfg := make([]model.SuperTroopSeasonCfg, 0)
	NacosConfigUnmarshal(data, &tmpSeasonsCfg)
	superTroopSeasons = tmpSeasonsCfg
	return myerr.SUCCESS
}

func SuperTroopSeasonSid() (int, int, int, *myerr.MyErr) {
	ts := util.Ntime()
	for _, cfg := range superTroopSeasons {
		if ts >= cfg.StartTime && (ts <= cfg.EndTime || cfg.EndTime == -1) {
			return cfg.SeasonId, cfg.StartTime, cfg.EndTime, myerr.SUCCESS
		}
	}
	return 0, 0, 0, myerr.LACK_OF_CONFIG
}
