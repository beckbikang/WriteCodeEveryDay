package service

import (
	"context"
	"errors"

	"github.com/beckbikang/WriteCodeEveryDay/internal/cache"
	"github.com/beckbikang/WriteCodeEveryDay/internal/model"
	"github.com/beckbikang/WriteCodeEveryDay/internal/mq"
	"github.com/beckbikang/WriteCodeEveryDay/pkg/app"
	"github.com/beckbikang/WriteCodeEveryDay/pkg/auth"
	"github.com/beckbikang/WriteCodeEveryDay/util"
)

type CreatePassportRequest struct {
	UserName string `form:"user_name" binding:"required,min=2,max=50"`
	Password string `form:"password" binding:"required,min=6,max=50"`
	Email    string `form:"email" binding:"required,email,max=100"`
	UserInfo string `form:"user_info" binding:"max=1024"`
	HeadImg  string `form:"head_img" binding:"max=1024"`
}

type UpdatePassportRequest struct {
	ID       uint64 `form:"id" binding:"required,gte=1"`
	UserName string `form:"user_name" binding:"min=2,max=50"`
	Email    string `form:"email" binding:"email,max=100"`
	UserInfo string `form:"user_info" binding:"max=1024"`
	HeadImg  string `form:"head_img" binding:"max=1024"`
}

type BlockPassportRequest struct {
	ID uint64 `form:"id" binding:"required,gte=1"`
}

type LoginRequest struct {
	UserName string `form:"user_name" binding:"required,min=2,max=50"`
	Password string `form:"password" binding:"required,min=6,max=50"`
}

func (svc *Service) CreatePassport(param *CreatePassportRequest) error {
	// 检查账号是否已存在
	exist, err := svc.ExistPassport(param.UserName)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("账号已存在")
	}

	// 密码加密
	password := util.GetMd5Str(param.Password)

	// 创建账号
	err = svc.dao.CreatePassport(param.UserName, password, param.Email, param.UserInfo, param.HeadImg)
	if err != nil {
		return err
	}

	// 发送用户注册消息到Kafka
	message := map[string]interface{}{
		"type": "user_registration",
		"data": map[string]string{
			"user_name": param.UserName,
			"email":     param.Email,
		},
	}
	return mq.SendMessage(message)
}

func (svc *Service) UpdatePassport(param *UpdatePassportRequest) error {
	// 更新数据库
	err := svc.dao.UpdatePassport(param.ID, param.UserName, param.Email, param.UserInfo, param.HeadImg)
	if err != nil {
		return err
	}

	// 删除缓存
	cache.DeletePassportCache(param.UserName)

	return nil
}

func (svc *Service) BlockPassport(param *BlockPassportRequest) error {
	// 获取用户信息
	passport, err := svc.dao.GetPassport(param.ID)
	if err != nil {
		return err
	}

	// 更新数据库
	err = svc.dao.BlockPassport(param.ID)
	if err != nil {
		return err
	}

	// 删除缓存
	cache.DeletePassportCache(passport.UserName)

	return nil
}

func (svc *Service) UnBlockPassport(param *BlockPassportRequest) error {
	// 获取用户信息
	passport, err := svc.dao.GetPassport(param.ID)
	if err != nil {
		return err
	}

	// 更新数据库
	err = svc.dao.UnBlockPassport(param.ID)
	if err != nil {
		return err
	}

	// 删除缓存
	cache.DeletePassportCache(passport.UserName)

	return nil
}

func (svc *Service) ExistPassport(userName string) (bool, error) {
	// 先从缓存中查询
	_, err := cache.GetPassportCache(userName)
	if err == nil {
		return true, nil
	}

	// 缓存中不存在，从数据库查询
	exist, err := svc.dao.ExistPassport(userName)
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (svc *Service) Login(param *LoginRequest) (bool, error) {
	// 密码加密
	password := util.GetMd5Str(param.Password)

	// 先从缓存中查询账号
	passport, err := cache.GetPassportCache(param.UserName)
	if err != nil {
		// 缓存中不存在，从数据库查询
		passport, err = svc.dao.GetPassportByUserNameAndPassword(param.UserName, password)
		if err != nil {
			return false, err
		}
	}

	// 验证密码
	if passport.UserPassword != password {
		return false, errors.New("密码错误")
	}

	// 检查账号状态
	if passport.Status == model.PassportStatusBlocked {
		return false, errors.New("账号已被拉黑")
	}

	// 设置缓存
	cache.SetPassportCache(param.UserName, &passport)

	return true, nil
}

func (svc *Service) GenerateToken(username string) (string, error) {
	return auth.GenerateToken(username)
}
