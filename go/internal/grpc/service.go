package grpc

import (
	"context"
	"errors"
	"log"

	"github.com/beckbikang/WriteCodeEveryDay/internal/model"
	"github.com/beckbikang/WriteCodeEveryDay/internal/service"
	"github.com/beckbikang/WriteCodeEveryDay/proto"
)

type PassportService struct {
	proto.UnimplementedPassportServiceServer
	svc *service.Service
}

func NewPassportService(svc *service.Service) *PassportService {
	return &PassportService{svc: svc}
}

func (s *PassportService) CreatePassport(ctx context.Context, req *proto.CreatePassportRequest) (*proto.CreatePassportResponse, error) {
	// 转换请求参数
	param := &service.CreatePassportRequest{
		UserName: req.UserName,
		Password: req.Password,
		Email:    req.Email,
		UserInfo: req.UserInfo,
		HeadImg:  req.HeadImg,
	}

	// 调用服务层方法
	err := s.svc.CreatePassport(param)
	if err != nil {
		return &proto.CreatePassportResponse{
			Code:    1,
			Message: err.Error(),
		}, nil
	}

	return &proto.CreatePassportResponse{
		Code:    0,
		Message: "success",
	}, nil
}

func (s *PassportService) UpdatePassport(ctx context.Context, req *proto.UpdatePassportRequest) (*proto.UpdatePassportResponse, error) {
	// 转换请求参数
	param := &service.UpdatePassportRequest{
		ID:       req.Id,
		UserName: req.UserName,
		Email:    req.Email,
		UserInfo: req.UserInfo,
		HeadImg:  req.HeadImg,
	}

	// 调用服务层方法
	err := s.svc.UpdatePassport(param)
	if err != nil {
		return &proto.UpdatePassportResponse{
			Code:    1,
			Message: err.Error(),
		}, nil
	}

	return &proto.UpdatePassportResponse{
		Code:    0,
		Message: "success",
	}, nil
}

func (s *PassportService) BlockPassport(ctx context.Context, req *proto.BlockPassportRequest) (*proto.BlockPassportResponse, error) {
	// 转换请求参数
	param := &service.BlockPassportRequest{
		ID: req.Id,
	}

	// 调用服务层方法
	err := s.svc.BlockPassport(param)
	if err != nil {
		return &proto.BlockPassportResponse{
			Code:    1,
			Message: err.Error(),
		}, nil
	}

	return &proto.BlockPassportResponse{
		Code:    0,
		Message: "success",
	}, nil
}

func (s *PassportService) UnBlockPassport(ctx context.Context, req *proto.UnBlockPassportRequest) (*proto.UnBlockPassportResponse, error) {
	// 转换请求参数
	param := &service.BlockPassportRequest{
		ID: req.Id,
	}

	// 调用服务层方法
	err := s.svc.UnBlockPassport(param)
	if err != nil {
		return &proto.UnBlockPassportResponse{
			Code:    1,
			Message: err.Error(),
		}, nil
	}

	return &proto.UnBlockPassportResponse{
		Code:    0,
		Message: "success",
	}, nil
}

func (s *PassportService) ExistPassport(ctx context.Context, req *proto.ExistPassportRequest) (*proto.ExistPassportResponse, error) {
	// 调用服务层方法
	exist, err := s.svc.ExistPassport(req.UserName)
	if err != nil {
		return &proto.ExistPassportResponse{
			Code:    1,
			Message: err.Error(),
			Exist:   false,
		}, nil
	}

	return &proto.ExistPassportResponse{
		Code:    0,
		Message: "success",
		Exist:   exist,
	}, nil
}

func (s *PassportService) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	// 转换请求参数
	param := &service.LoginRequest{
		UserName: req.UserName,
		Password: req.Password,
	}

	// 调用服务层方法
	success, err := s.svc.Login(param)
	if err != nil {
		return &proto.LoginResponse{
			Code:    1,
			Message: err.Error(),
			Success: false,
		}, nil
	}

	return &proto.LoginResponse{
		Code:    0,
		Message: "success",
		Success: success,
	}, nil
}
