package config

import (
	"math"

	"dm/dm/internal/model"
	"dm/pkg/util"
)

func GetInfWarSeasonId(ts int) (int, int) {
	if ts == 0 {
		ts = util.Ntime()
	}
	startTs, continueDay, intervalDay, unitSecond := GetModuleBasisMustData(model.InfWar)
	crossSecond := unitSecond * (continueDay + intervalDay)
	// 1赛季及以上
	nSeason := (ts-startTs)/crossSecond + model.InfWarSeason
	startTime := startTs + (nSeason-model.InfWarSeason)*crossSecond
	day := int(math.Floor(float64(ts-startTime) / float64(unitSecond)))

	return nSeason, day
}

func GetInfWarSeasonTime(sid int) (int, int) {
	// 获取赛季相关时间
	startTs, continueDay, intervalDay, unitSecond := GetModuleBasisMustData(model.InfWar)
	// 获取一个赛季总时间
	crossSecond := unitSecond * (continueDay + intervalDay)
	// 获取赛季开始时间
	startTime := startTs + (sid-model.InfWarSeason)*crossSecond
	// 获取赛季结束时间
	endTime := startTime + crossSecond - 1
	return startTime, endTime
}
