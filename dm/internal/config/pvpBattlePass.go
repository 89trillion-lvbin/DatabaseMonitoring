package config

import (
	"dm/dm/internal/model"
	"dm/pkg/util"
)

// 返回赛季开始和结束时间
func GetBattlePassSeasonInfo(sid int) (int, int) {
	var startTime, endTime int
	startTs, continueDay, intervalDay, unitSecond := GetModuleBasisMustData(model.PVP)

	crossSecond := (continueDay + intervalDay) * unitSecond
	// 20赛季及以上
	startTime = startTs + (sid-model.BpSeason)*crossSecond
	endTime = startTime + crossSecond - 1
	return startTime, endTime
}

// 返回当前bp赛季和开启结束时间
func GetBattlePassCurSeasonId() int {
	ts := util.Ntime()

	startTs, continueDay, intervalDay, unitSecond := GetModuleBasisMustData(model.PVP)
	// 20赛季及以上
	crossSecond := (continueDay + intervalDay) * unitSecond
	nSeason := (ts-startTs)/crossSecond + model.BpSeason
	return nSeason
}
