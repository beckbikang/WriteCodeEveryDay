package service

import (
	"context"

	"github.com/beckbikang/WriteCodeEveryDay/internal/dao"
	"github.com/jinzhu/gorm"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context, db *gorm.DB) *Service {
	svc := &Service{ctx: ctx}
	svc.dao = dao.New(db)
	return svc
}
