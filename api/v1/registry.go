package v1

import (
	"mall/api/v1/cart"
	"mall/api/v1/order"
	"mall/api/v1/product"
	"mall/api/v1/user"
)

type Registry struct {
	UserApi        user.IUserApi
	UserAddressApi user.IUserAddressApi
	ProductApi     product.IProductApi
	CartApi        cart.ICartApi
	OrderApi       order.IOrderApi
}

func (r *Registry) NewRegister() {
	r.UserApi = user.NewUserApi()
	r.UserAddressApi = user.NewUserAddressApi()
	r.ProductApi = product.NewProductApi()
	r.CartApi = cart.NewCartApi()
	r.OrderApi = order.NewProductApi()
}
