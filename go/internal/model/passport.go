package model

import (
	"github.com/jinzhu/gorm"
)

const (
	STATUS_NORMAL  = 1
	STATUS_BLOCKED = 0
)

type Passport struct {
	Model
	ID       uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	UserName string `gorm:"unique_index:uk_user_name;column:user_name;type:varchar(50)" json:"user_name"`
	Password string `gorm:"column:password;type:varchar(100)" json:"password"`
	Email    string `gorm:"unique_index:uk_email;column:user_email;type:varchar(100)" json:"user_email"`
	UserInfo string `gorm:"column:user_info;type:varchar(1024)" json:"user_info"`
	HeadImg  string `gorm:"column:head_img;type:varchar(1024)" json:"head_img"`
	Status   uint8  `gorm:"column:status;type:tinyint(1);default:1" json:"status"`
	CreatedOn  uint32 `gorm:"column:created_on;type:int(11);default:0" json:"created_on"`
	ModifiedOn uint32 `gorm:"column:modified_on;type:int(11);default:0" json:"modified_on"`
	DeletedOn  uint32 `gorm:"column:deleted_on;type:int(11);default:0" json:"deleted_on"`
	IsDel      uint8  `gorm:"column:is_del;type:tinyint(1);default:0" json:"is_del"`
}

func (p Passport) TableName() string {
	return "passport"
}

func (p Passport) Get(db *gorm.DB) (Passport, error) {
	var passport Passport
	db = db.Where("is_del = ?", 0)
	err := db.First(&passport).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return passport, err
	}

	return passport, nil
}

func (p Passport) GetByID(db *gorm.DB, id uint64) (Passport, error) {
	var passport Passport
	db = db.Where("id = ? AND is_del = ?", id, 0)
	err := db.First(&passport).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return passport, err
	}

	return passport, nil
}

func (p Passport) GetByUserName(db *gorm.DB, userName string) (Passport, error) {
	var passport Passport
	db = db.Where("user_name = ? AND is_del = ?", userName, 0)
	err := db.First(&passport).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return passport, err
	}

	return passport, nil
}

func (p Passport) GetByUserNameAndPassword(db *gorm.DB, userName, password string) (Passport, error) {
	var passport Passport
	db = db.Where("user_name = ? AND password = ? AND is_del = ?", userName, password, 0)
	err := db.First(&passport).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return passport, err
	}

	return passport, nil
}

func (p Passport) Exist(db *gorm.DB, userName string) (bool, error) {
	var passport Passport
	db = db.Where("user_name = ? AND is_del = ?", userName, 0)
	err := db.First(&passport).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if err == gorm.ErrRecordNotFound {
		return false, nil
	}

	return true, nil
}

func (p Passport) Create(db *gorm.DB) error {
	return db.Create(&p).Error
}

func (p Passport) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&p).Where("id = ? AND is_del = ?", p.ID, 0).Updates(values).Error
}

func (p Passport) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", p.Model.ID, 0).Delete(&p).Error
}

func (p Passport) List(db *gorm.DB, pageOffset, pageSize int) ([]*Passport, error) {
	var passports []*Passport
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	db = db.Where("is_del = ?", 0)
	if err = db.Find(&passports).Error; err != nil {
		return nil, err
	}

	return passports, nil
}
