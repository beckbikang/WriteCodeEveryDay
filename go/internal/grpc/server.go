package grpc

import (
	"context"
	"log"
	"net"

	"github.com/beckbikang/WriteCodeEveryDay/global"
	"github.com/beckbikang/WriteCodeEveryDay/internal/service"
	"github.com/beckbikang/WriteCodeEveryDay/proto"
	"google.golang.org/grpc"
)

func StartGrpcServer() {
	// 创建服务实例
	svc := service.New(context.Background(), global.DBEngine)
	passportService := NewPassportService(svc)

	// 创建gRPC服务器
	grpcServer := grpc.NewServer()
	proto.RegisterPassportServiceServer(grpcServer, passportService)

	// 监听端口
	lis, err := net.Listen("tcp", ":"+global.ServerSetting.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("gRPC server started on port %s", global.ServerSetting.GrpcPort)
	// 启动服务器
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
