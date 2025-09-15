package discovery

import (
	"context"
	"fmt"
	"time"

	"github.com/beckbikang/WriteCodeEveryDay/global"
	"go.etcd.io/etcd/clientv3"
)

var (
	EtcdClient *clientv3.Client
)

func SetupEtcd() error {
	config := clientv3.Config{
		Endpoints:   global.EtcdSetting.Endpoints,
		DialTimeout: time.Duration(global.EtcdSetting.DialTimeout) * time.Second,
	}
	
	if global.EtcdSetting.Username != "" && global.EtcdSetting.Password != "" {
		config.Username = global.EtcdSetting.Username
		config.Password = global.EtcdSetting.Password
	}
	
	var err error
	EtcdClient, err = clientv3.New(config)
	return err
}

func RegisterServiceEtcd(name, address string, port int) error {
	key := fmt.Sprintf("/services/%s/%s:%d", name, address, port)
	value := fmt.Sprintf("%s:%d", address, port)
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	_, err := EtcdClient.Put(ctx, key, value)
	return err
}

func DeregisterServiceEtcd(name, address string, port int) error {
	key := fmt.Sprintf("/services/%s/%s:%d", name, address, port)
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	_, err := EtcdClient.Delete(ctx, key)
	return err
}
