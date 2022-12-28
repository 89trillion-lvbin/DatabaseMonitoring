package model

import "dm/pkg/commonmodel"

// 超级大兵配置
type SuperTroopSeasonCfg struct {
	SeasonId     int                      `json:"seasonId"`
	StartTime    int                      `json:"startTime"`
	EndTime      int                      `json:"endTime"`
	Cond         SuperTroopCond           `json:"cond"`
	TotalRewards []SuperTroopSeasonReward `json:"totalRewards"`
	DailyRewards []SuperTroopDailyReward  `json:"dailyRewards"`
	DefLay       map[string]interface{}   `json:"defLay"`
	WallLay      map[string]interface{}   `json:"wallLay"`
	LvMap        map[string]interface{}   `json:"lvMap"`
	Theme        int                      `json:"theme"`
	Weather      int                      `json:"weather"`
	IslandId     int                      `json:"islandId"`
	LevelId      int                      `json:"levelId"`
	Troop        int                      `json:"troop"`
	DamageType   int                      `json:"damageType"`
	BanUnitText  string                   `json:"banUnitText"`
	Cvc          int                      `json:"cvc"` // 新增字段
}

type SuperTroopCond struct {
	Rarity   []int `json:"rarity"`
	MoveType []int `json:"moveType"`
}

type SuperTroopDailyReward struct {
	Level   int                         `json:"level"`
	Rewards []*commonmodel.RewardDetail `json:"rewards"`
}

type SuperTroopLayout struct {
	DefLay     map[int]int `json:"def_lay"`
	WallLay    map[int]int `json:"wall_lay"`
	LvMap      map[int]int `json:"lv_map"`
	Theme      int         `json:"theme"`
	Weather    int         `json:"weather"`
	IslandId   int         `json:"island_id"`
	LevelId    int         `json:"level_id"`
	DamageType int         `json:"damage_type"`
}

type SuperTroopSeasonReward struct {
	Start   int                         `json:"start"`
	End     int                         `json:"end"`
	Rewards []*commonmodel.RewardDetail `json:"rewards"`
}
