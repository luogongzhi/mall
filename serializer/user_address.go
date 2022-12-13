package serializer

import "mall/model"

type UserAddressVO struct {
	Id             uint64 `json:"id"`
	Name           string `json:"name"`
	Tel            string `json:"tel"`
	AddressDetails string `json:"address_details"`
}

type UserAddressCreateDTO struct {
	Name           string `json:"name" binding:"required"`
	Tel            string `json:"tel" binding:"required"`
	AddressDetails string `json:"address_details" binding:"required"`
}

type UserAddressUpdateDTO struct {
	Id             uint64 `json:"id" binding:"required"`
	Name           string `json:"name"`
	Tel            string `json:"tel"`
	AddressDetails string `json:"address_details"`
}

type UserAddressDeleteDTO struct {
	Id uint64 `json:"id" binding:"required"`
}

func NewUserAddressVO(userAddress *model.UserAddress) UserAddressVO {
	return UserAddressVO{
		Id:             userAddress.Id,
		Name:           userAddress.Name,
		Tel:            userAddress.Tel,
		AddressDetails: userAddress.AddressDetails,
	}
}

func NewUserAddressVOList(userAddressList []*model.UserAddress) (userAddressVOList []UserAddressVO) {
	for _, item := range userAddressList {
		userAddress := NewUserAddressVO(item)
		userAddressVOList = append(userAddressVOList, userAddress)
	}
	return userAddressVOList
}
