package global

import (
	"github.com/beckbikang/WriteCodeEveryDay/pkg/setting"
	"github.com/jinzhu/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	RedisSetting    *setting.RedisSettingS
	KafkaSetting    *setting.KafkaSettingS
	ConsulSetting   *setting.ConsulSettingS
	EtcdSetting     *setting.EtcdSettingS
	JWTSetting      *setting.AppSettingS
	DBEngine        *gorm.DB
)
