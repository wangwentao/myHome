package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

var (
	namingClient naming_client.INamingClient
	configClient config_client.IConfigClient
)

func initNacos() {
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
			//ClientConfig:  &clientConfig,
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

func main() {

	initNacos()

	services, err := namingClient.GetService(vo.GetServiceParam{
		//Clusters    []string `param:"clusters"`    //optional,default:DEFAULT
		ServiceName: "ping",
		// GroupName   string   `param:"groupName"`   //optional,default:DEFAULT_GROUP

	})

	fmt.Println(services, services.Hosts[0], err)

}
