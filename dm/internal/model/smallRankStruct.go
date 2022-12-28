package model

import "dm/pkg/commonmodel"

type SmallRankStruct struct {
	BasicSmallRankConfig
	Stages []SmallRankStage `json:"stages"`
}

type SmallRankCli struct {
	BasicSmallRankConfig
	Stage        int                     `json:"stage"`
	StageRewards []SmallRankStageRewards `json:"stageRewards"`
}

type SmallRankConfig struct {
	BasicSmallRankConfig
	Stages map[int]SmallRankStage `json:"stages"`
}

type BasicCondition struct {
	Type  int `json:"type"`
	Value int `json:"value"`
}

type BasicSmallRankConfig struct {
	Sid               int               `json:"sid"`
	StartTime         int               `json:"startTime"`
	EndTime           int               `json:"endTime"`
	CloseTime         int               `json:"closeTime"`
	RankType          int               `json:"rankType"`
	DisplayCondition  []BasicCondition  `json:"displayCondition"`
	Ext               BaseExt           `json:"ext"`
	ActivityCenterExt ActivityCenterExt `json:"activityCenterExt"`
}

type SmallRankStage struct {
	Stage        int                     `json:"stage"`
	ScoreStart   int                     `json:"scoreStart"`
	ScoreEnd     int                     `json:"scoreEnd"`
	StageRewards []SmallRankStageRewards `json:"stageRewards"`
}

type SmallRankStageRewards struct {
	Start   int                         `json:"start"`
	End     int                         `json:"end"`
	Rewards []*commonmodel.RewardDetail `json:"rewards"`
}

type BaseExt struct {
	Style         []Style `json:"style"`
	EntranceClose int     `json:"entranceClose"`
}

type ActivityCenterExt struct {
	Style Style `json:"style"`
}

type Style struct {
	Type        int      `json:"type"`
	Sort        int      `json:"sort"`
	ModelId     int      `json:"modelId"`
	JumpType    int      `json:"jumpType"`
	FormatType  int      `json:"formatType"`
	Sprite      []string `json:"sprite"`
	LanguageKey []string `json:"languageKey"`
	SpriteUrl   []string `json:"spriteUrl,omitempty"`
}
