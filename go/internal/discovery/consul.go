package discovery

import (
	"fmt"

	"github.com/beckbikang/WriteCodeEveryDay/global"
	"github.com/hashicorp/consul/api"
)

var (
	ConsulClient *api.Client
)

func SetupConsul() error {
	config := api.DefaultConfig()
	config.Address = global.ConsulSetting.Addr
	config.Scheme = global.ConsulSetting.Scheme
	
	var err error
	ConsulClient, err = api.NewClient(config)
	return err
}

func RegisterService(name, address string, port int) error {
	registration := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", name, address, port),
		Name:    name,
		Address: address,
		Port:    port,
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%d/health", address, port),
			Interval: "10s",
			Timeout:  "5s",
		},
	}
	
	return ConsulClient.Agent().ServiceRegister(registration)
}

func DeregisterService(serviceID string) error {
	return ConsulClient.Agent().ServiceDeregister(serviceID)
}
