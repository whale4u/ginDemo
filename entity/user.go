package entity

import (
	orm "ginDemo/common"
)

type User struct {
	id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

var Users []User

func (user User) Add() (id int64, err error) {
	result := orm.DB.Create(&user)
	id = user.id
	if result.Error != nil {
		err = result.Error
	}
	return
}

func (user *User) Delete(id int64) (Result User, err error) {
	user.id = id
	if err = orm.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return
	}
	Result = *user
	return
}

func (user *User) Inquire(name string) (Result User) {
	user.Username = name
	if err := orm.DB.First(&user, "username = ?", name).Error; err != nil {
		return
	}
	Result = *user
	return
}

func (user *User) Change(name, passwd string) (Result bool) {
	if err := orm.DB.Model(&user).Where("username = ?", name).Update("password", passwd); err != nil {
		Result = true
	} else {
		Result = false
	}
	return Result
}

func (user *User) Users() (AllUsers []User) {
	var users []User
	if err := orm.DB.Find(&users).Error; err != nil {
		return
	}
	return users
}
