package config

import (
	"dm/dm/client/nacosclient"
	"dm/dm/setting"
	"dm/pkg/myerr"
	"dm/pkg/util"

	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

const (
	SUPER_TROOP_Rewards = "super_troop_rewards"
	MODULE_BASIC        = "module_basic"
	SMALL_RANK          = "small_rank"
)

var refreshConfigMap = map[string]func(data, md5 string) *myerr.MyErr{
	SUPER_TROOP_Rewards: RefreshSuperTroopRewards,
	MODULE_BASIC:        RefreshModuleBasicData,
	SMALL_RANK:          RefreshSmallRankCfg,
}

func NacosConfigSetUp() {
	for _, item := range InitConfig {
		_, ok := refreshConfigMap[item.DataId]
		if !ok {
			continue
		}
		myErr := callBackListen(item.Content, refreshConfigMap[item.DataId])
		if myErr != myerr.SUCCESS {
			panic(myErr)
		}
	}

	for k := range refreshConfigMap {
		_ = nacosclient.NacosConfigClient.ListenConfig(vo.ConfigParam{
			DataId: k,
			Group:  setting.NacosSetting.ConfigGroupId,
			OnChange: func(namespace, group, dataId, data string) {
				_ = callBackListen(data, refreshConfigMap[dataId])
			},
		})
	}
}

func callBackListen(data string, f func(data, md5 string) *myerr.MyErr) *myerr.MyErr {
	md5 := util.EncodeMD5(data)
	return f(data, md5)
}
