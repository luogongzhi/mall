package model

type CartProduct struct {
	Model
	CartId    uint64
	ProductId uint64
	Total     uint16
}

func (CartProduct) TableName() string {
	return "cart_product"
}
