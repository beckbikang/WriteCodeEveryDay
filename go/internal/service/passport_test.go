package service

import (
	"context"
	"testing"

	"github.com/go-programming-tour-book/WriteCodeEveryDay/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestCreatePassport(t *testing.T) {
	// 创建测试服务实例
	svc := New(context.Background(), nil)

	// 测试参数
	param := &CreatePassportRequest{
		UserName: "testuser",
		Password: "password123",
		Email:    "test@example.com",
		UserInfo: "test user info",
		HeadImg:  "http://example.com/avatar.jpg",
	}

	// 测试创建账号
	err := svc.CreatePassport(param)
	assert.NoError(t, err)
}

func TestUpdatePassport(t *testing.T) {
	// 创建测试服务实例
	svc := New(context.Background(), nil)

	// 测试参数
	param := &UpdatePassportRequest{
		ID:       1,
		UserName: "updateduser",
		Email:    "updated@example.com",
		UserInfo: "updated user info",
		HeadImg:  "http://example.com/updated_avatar.jpg",
	}

	// 测试更新账号
	err := svc.UpdatePassport(param)
	assert.NoError(t, err)
}

func TestBlockAndUnblockPassport(t *testing.T) {
	// 创建测试服务实例
	svc := New(context.Background(), nil)

	// 测试拉黑账号
	blockParam := &BlockPassportRequest{ID: 1}
	err := svc.BlockPassport(blockParam)
	assert.NoError(t, err)

	// 测试解拉黑账号
	err = svc.UnBlockPassport(blockParam)
	assert.NoError(t, err)
}

func TestExistPassport(t *testing.T) {
	// 创建测试服务实例
	svc := New(context.Background(), nil)

	// 测试账号存在性检查
	exist, err := svc.ExistPassport("testuser")
	assert.NoError(t, err)
	assert.True(t, exist)
}

func TestLogin(t *testing.T) {
	// 创建测试服务实例
	svc := New(context.Background(), nil)

	// 测试登录参数
	param := &LoginRequest{
		UserName: "testuser",
		Password: "password123",
	}

	// 测试登录功能
	success, err := svc.Login(param)
	assert.NoError(t, err)
	assert.True(t, success)
}
