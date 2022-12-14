package serializer

import (
	"mall/model"
)

type CartCreateDeleteDTO struct {
	ProductId uint64 `json:"product_id" binding:"required"`
	Total     uint16 `json:"total" binding:"required""`
}

type CartVO struct {
	CartId        uint64          `json:"cart_id"`
	UserId        uint64          `json:"user_id"`
	Total         uint16          `json:"total"`
	CartProductVO []CartProductVO `json:"cart_product"`
}

func NewCartVO(cart *model.Cart, cartProductVOList []CartProductVO) CartVO {
	return CartVO{
		CartId:        cart.Id,
		UserId:        cart.UserId,
		Total:         cart.Total,
		CartProductVO: cartProductVOList,
	}
}
