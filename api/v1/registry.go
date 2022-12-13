package v1

import (
	"mall/api/v1/user"
)

type Registry struct {
	UserApi user.IUserApi
}

func (r *Registry) NewRegister() {
	r.UserApi = user.NewUserApi()
}
