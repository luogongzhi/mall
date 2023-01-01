package service

import (
	"context"
	"mall/dao"
	"mall/model"
	"mall/pkg/e"
	"mall/serializer"
	"net/http"
)

type UserAddressService struct{}

// List 根据用户Id查询用户地址信息
func (*UserAddressService) List(ctx context.Context, id uint64) serializer.ResponseResult {
	db := dao.NewDBClient(ctx)
	userAddressDao := dao.NewUserAddressDao(db)

	userAddress, exist, err := userAddressDao.GetByUserId(id)
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}
	if !exist {
		return serializer.ResponseResult{
			Code: http.StatusOK,
			Msg:  e.GetMsg(http.StatusOK),
		}
	}

	return serializer.ResponseResult{
		Code: http.StatusOK,
		Msg:  e.GetMsg(http.StatusOK),
		Data: map[string]interface{}{
			"userAddressList": serializer.NewUserAddressVOList(userAddress),
		},
	}
}

// Create 用户地址信息添加
func (*UserAddressService) Create(ctx context.Context, dto serializer.UserAddressCreateDTO, id uint64) serializer.ResponseResult {
	db := dao.NewDBClient(ctx)
	userAddressDao := dao.NewUserAddressDao(db)

	// 创建用户地址信息
	err := userAddressDao.Create(&model.UserAddress{
		UserId:         id,
		Name:           dto.Name,
		Tel:            dto.Tel,
		AddressDetails: dto.AddressDetails,
	})
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}

	return serializer.ResponseResult{
		Code: http.StatusOK,
		Msg:  e.GetMsg(http.StatusOK),
	}
}

// Delete 用户地址信息删除
func (*UserAddressService) Delete(ctx context.Context, userAddressId uint64, userId uint64) serializer.ResponseResult {
	db := dao.NewDBClient(ctx)
	userAddressDao := dao.NewUserAddressDao(db)

	// 删除用户地址信息
	row, err := userAddressDao.DeleteById(userId, userAddressId)
	if row == 0 {
		return serializer.ResponseResult{
			Code: e.ErrorNotExistUserAddress,
			Msg:  e.GetMsg(e.ErrorNotExistUserAddress),
		}
	}
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}

	return serializer.ResponseResult{
		Code: http.StatusOK,
		Msg:  e.GetMsg(http.StatusOK),
	}
}

// Update 修改用户地址信息
func (*UserAddressService) Update(ctx context.Context, dto serializer.UserAddressUpdateDTO, id uint64) serializer.ResponseResult {
	db := dao.NewDBClient(ctx)
	userAddressDao := dao.NewUserAddressDao(db)

	// 根据Id修改用户信息
	row, err := userAddressDao.UpdateById(dto.Id, &model.UserAddress{
		UserId:         id,
		Name:           dto.Name,
		Tel:            dto.Tel,
		AddressDetails: dto.AddressDetails,
	})
	if row == 0 {
		return serializer.ResponseResult{
			Code: e.ErrorNotExistUserAddress,
			Msg:  e.GetMsg(e.ErrorNotExistUserAddress),
		}
	}
	if err != nil {
		return serializer.ResponseResult{
			Code: e.ErrorDatabase,
			Msg:  e.GetMsg(e.ErrorDatabase),
		}
	}

	return serializer.ResponseResult{
		Code: http.StatusOK,
		Msg:  e.GetMsg(http.StatusOK),
	}
}
