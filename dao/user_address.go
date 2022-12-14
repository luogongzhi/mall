package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
)

type UserAddress struct {
	*gorm.DB
}

func NewUserAddressDao(ctx context.Context) *UserAddress {
	return &UserAddress{NewDBClient(ctx)}
}

// GetByUserId 根据用户id获取用户地址信息
func (dao *UserAddress) GetByUserId(userId uint64) (userAddress []*model.UserAddress, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.UserAddress{}).Where("user_id = ?", userId).Find(&userAddress).Count(&count).Error
	if count == 0 || err != nil {
		return nil, false, err
	}
	return userAddress, true, nil
}

// UpdateById 更新用户地址信息
func (dao *UserAddress) UpdateById(userAddressId uint64, userAddress *model.UserAddress) (row int64, err error) {
	result := dao.DB.Where("id = ? AND user_id = ?", userAddressId, userAddress.UserId).Updates(&userAddress)
	return result.RowsAffected, result.Error
}

// DeleteById 删除用户地址信息
func (dao *UserAddress) DeleteById(userId uint64, userAddressId uint64) (row int64, err error) {
	result := dao.DB.Where("user_id = ? AND id = ?", userId, userAddressId).Delete(&model.UserAddress{})
	return result.RowsAffected, result.Error
}

// Create 创建用户地址
func (dao *UserAddress) Create(userAddress *model.UserAddress) (err error) {
	return dao.DB.Create(&userAddress).Error
}
