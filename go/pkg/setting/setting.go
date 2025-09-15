package setting

import (
	"time"

	"github.com/spf13/viper"
)

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	GrpcPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	PageSize  int
	JwtSecret string
	JwtExpire int
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type RedisSettingS struct {
	Addr         string
	Password     string
	DB           int
	PoolSize     int
	MinIdleConns int
}

type KafkaSettingS struct {
	Brokers []string
	Topic   string
	Group   string
}

type ConsulSettingS struct {
	Addr string
	Scheme string
	Datacenter string
	Token string
}

type EtcdSettingS struct {
	Endpoints []string
	DialTimeout int
	Username string
	Password string
}

type Setting struct {
	v *viper.Viper
}

func NewSetting(paths ...string) (*Setting, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	for _, path := range paths {
		v.AddConfigPath(path)
	}
	v.SetConfigName("config")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Setting{v: v}, nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.v.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}

func (s *Setting) ReadRedisSection(k string, v *RedisSettingS) error {
	err := s.v.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}

func (s *Setting) ReadKafkaSection(k string, v *KafkaSettingS) error {
	err := s.v.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}

func (s *Setting) ReadConsulSection(k string, v *ConsulSettingS) error {
	err := s.v.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}

func (s *Setting) ReadEtcdSection(k string, v *EtcdSettingS) error {
	err := s.v.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
