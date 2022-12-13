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

func NewUserAddressDaoByDB(db *gorm.DB) *UserAddress {
	return &UserAddress{db}
}

// GetUserAddressByUserId 根据用户id获取用户地址信息
func (dao *UserAddress) GetUserAddressByUserId(userId uint64) (userAddress []*model.UserAddress, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.UserAddress{}).Where("user_id = ?", userId).Find(&userAddress).Count(&count).Error
	if count == 0 || err != nil {
		return nil, false, err
	}
	return userAddress, true, nil
}

// UpdateUserAddressById 更新用户地址信息
func (dao *UserAddress) UpdateUserAddressById(userAddressId uint64, userAddress *model.UserAddress) (row int64, err error) {
	result := dao.DB.Where("id = ? AND user_id = ?", userAddressId, userAddress.UserId).Updates(&userAddress)
	return result.RowsAffected, result.Error
}

// DeleteUserAddressById 删除用户地址信息
func (dao *UserAddress) DeleteUserAddressById(userId uint64, userAddressId uint64) (row int64, err error) {
	result := dao.DB.Where("user_id = ? AND id = ?", userId, userAddressId).Delete(&model.UserAddress{})
	return result.RowsAffected, result.Error
}

// CreateUserAddress 创建用户地址
func (dao *UserAddress) CreateUserAddress(userAddress *model.UserAddress) (err error) {
	return dao.DB.Create(&userAddress).Error
}
