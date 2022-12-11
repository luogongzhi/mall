package model

type Order struct {
	Model
	UserId         uint64
	AddressName    string
	AddressTel     string
	AddressDetails string
	ProductAmount  float64
	FreightAmount  float64
	TotalAmount    float64
	Status         bool
}

func (Order) TableName() string {
	return "order"
}
