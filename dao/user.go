package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

// GetUserById 根据id获取用户
func (dao *UserDao) GetUserById(id uint64) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("id = ?", id).Find(&user).Count(&count).Error
	if count == 0 || err != nil {
		return nil, false, err
	}
	return user, true, nil
}

// UpdateUserById 更新用户信息
func (dao *UserDao) UpdateUserById(id uint64, user *model.User) error {
	return dao.DB.Where("id = ?", id).Updates(&user).Error
}

// ExistOrNotByUserName 判断是否存在该名字
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("username = ?", userName).Find(&user).Count(&count).Error
	if count == 0 || err != nil {
		return nil, false, err
	}
	return user, true, nil
}

// CreateUser 创建用户
func (dao *UserDao) CreateUser(user *model.User) (err error) {
	return dao.DB.Create(&user).Error
}
