package model

type OrderProduct struct {
	Model
	OrderId   uint64
	ProductId uint64
	Total     uint16
}

func (OrderProduct) TableName() string {
	return "order_product"
}
