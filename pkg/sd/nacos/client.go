package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

// Client is a wrapper around the Nacos API.
type Client interface {

	// INamingClient Service Client*/
	naming_client.INamingClient

	// IConfigClient Config Client*/
	config_client.IConfigClient
}

type client struct {
	nc naming_client.INamingClient
	cc config_client.IConfigClient
}

func (c *client) CloseClient() {
	if c.nc != nil {
		c.nc.CloseClient()
	}
	if c.cc != nil {
		c.cc.CloseClient()
	}
}

func NewCLient(param *vo.NacosClientParam) Client {
	namingClient, errN := clients.NewNamingClient(*param)
	if errN != nil {
		fmt.Println(errN)
	}
	configClient, errC := clients.NewConfigClient(*param)
	if errC != nil {
		fmt.Println(errN)
	}

	return &client{nc: namingClient, cc: configClient}
}

/** Naming Client **/

func (c *client) RegisterInstance(param vo.RegisterInstanceParam) (bool, error) {

	return c.nc.RegisterInstance(param)
}

func (c *client) DeregisterInstance(param vo.DeregisterInstanceParam) (bool, error) {

	return c.nc.DeregisterInstance(param)
}

func (c *client) UpdateInstance(param vo.UpdateInstanceParam) (bool, error) {

	return c.nc.UpdateInstance(param)
}

func (c *client) GetService(param vo.GetServiceParam) (model.Service, error) {

	return c.nc.GetService(param)
}

func (c *client) SelectAllInstances(param vo.SelectAllInstancesParam) ([]model.Instance, error) {

	return c.nc.SelectAllInstances(param)
}

func (c *client) SelectInstances(param vo.SelectInstancesParam) ([]model.Instance, error) {

	return c.nc.SelectInstances(param)
}

func (c *client) SelectOneHealthyInstance(param vo.SelectOneHealthInstanceParam) (*model.Instance, error) {

	return c.nc.SelectOneHealthyInstance(param)
}

func (c *client) Subscribe(param *vo.SubscribeParam) error {

	return c.nc.Subscribe(param)
}

func (c *client) Unsubscribe(param *vo.SubscribeParam) error {

	return c.nc.Unsubscribe(param)
}

func (c *client) GetAllServicesInfo(param vo.GetAllServiceInfoParam) (model.ServiceList, error) {

	return c.nc.GetAllServicesInfo(param)
}

/*Config Client*/

func (c *client) GetConfig(param vo.ConfigParam) (string, error) {

	return c.cc.GetConfig(param)
}

func (c *client) PublishConfig(param vo.ConfigParam) (bool, error) {

	return c.cc.PublishConfig(param)
}

func (c *client) DeleteConfig(param vo.ConfigParam) (bool, error) {

	return c.cc.DeleteConfig(param)
}

func (c *client) ListenConfig(params vo.ConfigParam) (err error) {

	return c.cc.ListenConfig(params)
}

func (c *client) CancelListenConfig(params vo.ConfigParam) (err error) {

	return c.cc.CancelListenConfig(params)
}

func (c *client) SearchConfig(param vo.SearchConfigParm) (*model.ConfigPage, error) {

	return c.cc.SearchConfig(param)
}
