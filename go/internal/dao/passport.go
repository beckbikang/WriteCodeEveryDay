package dao

import (
	"github.com/beckbikang/WriteCodeEveryDay/internal/model"
)

func (d *Dao) GetPassport(id uint64) (model.Passport, error) {
	passport := model.Passport{Model: model.Model{ID: id}}
	return passport.Get(d.engine)
}

func (d *Dao) GetPassportByUserName(userName string) (model.Passport, error) {
	passport := model.Passport{UserName: userName}
	return passport.GetByUserName(d.engine, userName)
}

func (d *Dao) GetPassportByUserNameAndPassword(userName, password string) (model.Passport, error) {
	passport := model.Passport{}
	return passport.GetByUserNameAndPassword(d.engine, userName, password)
}

func (d *Dao) ExistPassport(userName string) (bool, error) {
	passport := model.Passport{UserName: userName}
	return passport.Exist(d.engine, userName)
}

func (d *Dao) CreatePassport(userName, password, email, userInfo, headImg string) error {
	passport := model.Passport{
		UserName:     userName,
		UserPassword: password,
		UserEmail:    email,
		UserInfo:     userInfo,
		HeadImg:      headImg,
		Status:       model.STATUS_NORMAL,
	}
	return passport.Create(d.engine)
}

func (d *Dao) UpdatePassport(id uint64, userName, email, userInfo, headImg string) error {
	passport := model.Passport{Model: model.Model{ID: id}}
	values := map[string]interface{}{
		"modified_by": "system",
	}
	if userName != "" {
		values["user_name"] = userName
	}
	if email != "" {
		values["user_email"] = email
	}
	if userInfo != "" {
		values["user_info"] = userInfo
	}
	if headImg != "" {
		values["head_img"] = headImg
	}
	return passport.Update(d.engine, values)
}

func (d *Dao) BlockPassport(id uint64) error {
	passport := model.Passport{Model: model.Model{ID: id}}
	values := map[string]interface{}{
		"status":      model.STATUS_BLOCKED,
		"modified_by": "system",
	}
	return passport.Update(d.engine, values)
}

func (d *Dao) UnBlockPassport(id uint64) error {
	passport := model.Passport{Model: model.Model{ID: id}}
	values := map[string]interface{}{
		"status":      model.STATUS_NORMAL,
		"modified_by": "system",
	}
	return passport.Update(d.engine, values)
}
