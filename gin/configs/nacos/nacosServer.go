package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var (
	namingClient naming_client.INamingClient
	configClient config_client.IConfigClient
)

func RegisterNacosService() {

	configNcosClient()

	registerService()
}

func DeregisterNacosService() {

	success, err := namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          "localhost",
		Port:        8880,
		ServiceName: "ping",
		Ephemeral:   true,
		//Cluster:     "cluster-a", // default value is DEFAULT
		//GroupName:   "group-a",   // default value is DEFAULT_GROUP
	})
	fmt.Println(success, err)
}

func registerService() {
	success, errR := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "localhost",
		Port:        8880,
		ServiceName: "ping",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"path": "/home/ping"},
		//ClusterName: "cluster-a", // 默认值DEFAULT
		//GroupName:   "group-a",   // 默认值DEFAULT_GROUP
	})
	fmt.Println(success, errR)
}

func configNcosClient() {
	//create clientConfig
	clientConfig := constant.ClientConfig{
		// NamespaceId:         "nacos_demo001", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./nacos/log",
		CacheDir:            "./nacos/cache",
		LogLevel:            "debug",
	}

	// 至少一个ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "localhost",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}

	var err error
	namingClient, err = clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	fmt.Println(err)

	configClient, err = clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	fmt.Println(err)
}
