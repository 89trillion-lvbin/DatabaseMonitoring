package config

import (
	"dm/dm/internal/model"
	"dm/pkg/util"
)

func GetPointRaceSid(isElite bool) int {
	moduleType := model.Elite
	if !isElite {
		moduleType = model.Master
	}
	startTs, continueDay, intervalDay, unitSecond := GetModuleBasisMustData(moduleType)

	crossSecond := (continueDay + intervalDay) * unitSecond

	ts := util.Ntime()
	sid := (ts-startTs)/(crossSecond) + model.PtSeason

	return sid
}

func GetPointRaceStartEndTime(sid int, isElite bool) (int, int, int) {
	moduleType := model.Elite
	if !isElite {
		moduleType = model.Master
	}
	startTs, continueDay, intervalDay, unitSecond := GetModuleBasisMustData(moduleType)

	crossSecond := (continueDay + intervalDay) * unitSecond
	startTime := startTs + (sid-model.PtSeason)*crossSecond
	attackEnd := startTime + (continueDay * unitSecond)
	endTime := startTime + crossSecond
	return startTime, attackEnd, endTime
}
