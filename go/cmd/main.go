package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/beckbikang/WriteCodeEveryDay/global"
	"github.com/beckbikang/WriteCodeEveryDay/internal/cache"
	"github.com/beckbikang/WriteCodeEveryDay/internal/discovery"
	"github.com/beckbikang/WriteCodeEveryDay/internal/model"
	"github.com/beckbikang/WriteCodeEveryDay/internal/mq"
	"github.com/beckbikang/WriteCodeEveryDay/internal/routers"
	"github.com/beckbikang/WriteCodeEveryDay/pkg/setting"
)

func init() {
	// 初始化配置
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	// 初始化数据库
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	// 初始化Redis缓存
	err = cache.SetupRedis()
	if err != nil {
		log.Fatalf("init.setupRedis err: %v", err)
	}

	// 初始化Consul
	err = discovery.SetupConsul()
	if err != nil {
		log.Fatalf("init.setupConsul err: %v", err)
	}

	// 初始化Etcd
	err = discovery.SetupEtcd()
	if err != nil {
		log.Fatalf("init.setupEtcd err: %v", err)
	}

	// 初始化Kafka生产者
	err = mq.SetupKafkaProducer()
	if err != nil {
		log.Fatalf("init.setupKafkaProducer err: %v", err)
	}

	// 初始化Kafka消费者
	err = mq.SetupKafkaConsumer()
	if err != nil {
		log.Fatalf("init.setupKafkaConsumer err: %v", err)
	}
}

func main() {
	// 设置Gin运行模式
	gin.SetMode(global.ServerSetting.RunMode)

	// 创建路由
	router := routers.NewRouter()

	// 启动gRPC服务器
	go grpc.StartGrpcServer()

	// 启动Kafka消费者
	go mq.ConsumeMessages(func(message []byte) error {
		log.Printf("Received message: %s", string(message))
		// 处理用户注册消息
		return nil
	})

	// 创建HTTP服务器
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// 启动服务器
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err: %v", err)
		}
	}()

	// 优雅关闭
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// 关闭Kafka连接
	mq.CloseKafkaProducer()
	mq.CloseKafkaConsumer()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func setupSetting() error {
	s, err := setting.NewSetting("configs")
	if err != nil {
		return err
	}

	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Redis", &global.RedisSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Kafka", &global.KafkaSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Consul", &global.ConsulSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Etcd", &global.EtcdSetting)
	if err != nil {
		return err
	}

	// 初始化JWT设置，使用App配置中的JWT相关参数
	global.JWTSetting = global.AppSetting

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupRedis() error {
	// 这里可以初始化Redis连接
	// 暂时留空，后续实现
	return nil
}
