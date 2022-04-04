package model

import (
	"tf-sso/util/common"

	"github.com/labstack/echo/v4"
)

type User struct {
	Owner     string `xorm:"varchar(100) notnull pk" json:"owner"`
	Username  string `xorm:"varchar(100) notnull pk" json:"username"`
	CreatedAt int64  `xorm:"created" json:"created"`
	UpdatedAt int64  `xorm:"updated" json:"updated"`

	ID            string `xorm:"varchar(100) index" json:"id" param:"id" validate:"required"` // 用户id
	Type          string `xorm:"varchar(100)" json:"type"`                                    // 用户类型
	Password      string `xorm:"varchar(100)" json:"password"`                                // 密码
	PasswordSalt  string `xorm:"varchar(100)" json:"passwordSalt"`                            // 密码盐
	DisplayName   string `xorm:"varchar(100)" json:"displayName"`                             // 用户显示名称
	Avatar        string `xorm:"varchar(500)" json:"avatar"`                                  // 头像
	Email         string `xorm:"varchar(100) index" json:"email"`                             // 电子邮件
	Phone         string `xorm:"varchar(100) index" json:"phone"`                             // 电话
	Location      string `xorm:"varchar(100)" json:"location"`                                // 地址
	Bio           string `xorm:"varchar(100)" json:"bio"`                                     // 描述
	Region        string `xorm:"varchar(100)" json:"region"`
	Gender        string `xorm:"varchar(100)" json:"gender"`
	IsOnline      bool   `json:"isOnline"`
	IsAdmin       bool   `json:"isAdmin"`
	IsGlobalAdmin bool   `json:"isGlobalAdmin"`
	IsForbidden   bool   `json:"isForbidden"`
	IsDeleted     bool   `json:"isDeleted"`

	LastSigninTime string `xorm:"varchar(100)" json:"lastSigninTime"`
	LastSigninIp   string `xorm:"varchar(100)" json:"lastSigninIp"`
}

func (u *User) Add() error {
	_, err := Engine.Insert(u)
	return err
}

func (u *User) Delete() error {
	_, err := Engine.Delete(u)
	return err
}

func (u *User) Get() error {
	_, err := Engine.Get(u)
	return err
}

func (u *User) Update() error {
	_, err := Engine.ID(u.ID).Update(u)
	return err
}

func (u *User) Count() (int64, error) {
	total, err := Engine.Count(u)
	return total, err
}

func GetUsers(c echo.Context) (*[]User, error) {
	pageSize, offset := common.Paginate(c)
	users := make([]User, 0, pageSize)
	err := Engine.Limit(pageSize, offset).Find(&users)
	return &users, err
}
