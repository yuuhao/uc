package utils

import (
	"fmt"
	"uc/config"

	"github.com/ghodss/yaml"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var naocsConf = naocs{}

type naocs struct {
	client config_client.IConfigClient
}

func (n *naocs) setConfig() {
	content, _ := n.client.GetConfig(vo.ConfigParam{
		DataId: "uc",
		Group:  "study",
	})

	conf := &config.ConfigMap{}

	if err := yaml.Unmarshal([]byte(content), conf); err != nil {
		panic(fmt.Sprintf("nacos get fail \n %v", err))
	}

	config.Cfg = conf
}

func InitNacos() {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848, constant.WithContextPath("/nacos")),
	}

	//create ClientConfig
	clientCfg := &constant.ClientConfig{
		NamespaceId: "c5da25b1-1aac-4f1b-b3f9-b1bba9342ca1",
		TimeoutMs:   5000,
		LogLevel:    "error",
		LogDir:      "D:/tmp/nacos/log",
		CacheDir:    "D:/tmp/nacos/cache",
		//NotLoadCacheAtStart: true,
		//RotateTime:          "1h",
		//MaxAge:              3,
	}

	// create config client
	var err error
	naocsConf.client, err = clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  clientCfg,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}
	_, err = naocsConf.client.GetConfig(vo.ConfigParam{
		DataId: "uc",
		Group:  "study",
	})

	if err != nil {
		panic(fmt.Sprintf("nacos get fail \n %v", err))
	}
	naocsConf.setConfig()
	go func() {
		err = naocsConf.client.ListenConfig(vo.ConfigParam{
			DataId: "uc",
			Group:  "study",
			OnChange: func(namespace, group, dataId, data string) {
				naocsConf.setConfig()
			},
		})
		if err != nil {
			panic(fmt.Sprintf("nacos get fail \n %v", err))
		}
	}()
}
