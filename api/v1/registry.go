package v1

import (
	"mall/api/v1/user"
)

type Registry struct {
	UserApi        user.IUserApi
	UserAddressApi user.IUserAddressApi
}

func (r *Registry) NewRegister() {
	r.UserApi = user.NewUserApi()
	r.UserAddressApi = user.NewUserAddressApi()
}
