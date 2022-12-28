package nacosclient

import (
	"dm/dm/setting"
	"dm/pkg/nacos"

	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
)

var NacosConfigClient config_client.IConfigClient

func SetUp() {
	NacosConfigClient = nacos.NewNacosConfigClient(setting.NacosClientConfig, setting.NacosServerConfigs)
}
