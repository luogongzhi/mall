package serializer

type OrderCreateDTO struct {
	AddressName    string `json:"address_name" binding:"required"`
	AddressTel     string `json:"address_tel" binding:"required"`
	AddressDetails string `json:"address_details" binding:"required"`
}

type OrderUpdateDTO struct {
	OrderId        uint64 `json:"order_id" binding:"required"`
	AddressName    string `json:"address_name"`
	AddressTel     string `json:"address_tel"`
	AddressDetails string `json:"address_details"`
	Status         bool   `json:"status"`
}

type OrderDeleteDTO struct {
	Id uint64 `json:"id" binding:"required"`
}

type OrderVO struct {
	OrderId        uint64 `json:"order_id"`
	UserId         uint64
	AddressName    string           `json:"address_name"`
	AddressTel     string           `json:"address_tel"`
	AddressDetails string           `json:"address_details"`
	OrderProductVO []OrderProductVO `json:"order_product"`
	ProductAmount  float64          `json:"product_amount"`
	FreightAmount  float64          `json:"freight_amount"`
	TotalAmount    float64          `json:"total_amount"`
	Status         string           `json:"status"`
}
