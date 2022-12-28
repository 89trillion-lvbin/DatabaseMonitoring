package config

import (
	"dm/dm/client/nacosclient"
	"dm/dm/setting"
	"dm/pkg/myerr"

	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

var configMap = map[string]func(data string) *myerr.MyErr{}

var InitConfig []model.ConfigItem

func InitConfigs() {
	getConfig()
}

func getConfig() {
	searchConfig := make([]model.ConfigItem, 0)
	for i := 1; i < 20; i++ {
		searchPage, err := nacosclient.NacosConfigClient.SearchConfig(vo.SearchConfigParam{
			Search:   "accurate",
			DataId:   "",
			Group:    setting.NacosSetting.ConfigGroupId,
			PageNo:   i,
			PageSize: 50, // 查询一个空间下最多的200条
		})
		if err != nil || searchPage == nil {
			panic(err)
		}
		searchConfig = append(searchConfig, searchPage.PageItems...)
		if len(searchPage.PageItems) < 50 {
			break
		}
	}
	initAndListenConfig(searchConfig)
}

func initAndListenConfig(searchConfig []model.ConfigItem) {
	for _, v := range searchConfig {
		f, ok := configMap[v.DataId]
		if !ok {
			InitConfig = append(InitConfig, v)
			continue
		}
		myErr := f(v.Content)
		if myErr != myerr.SUCCESS {
			panic(myErr)
		}
	}
}
