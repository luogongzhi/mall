package model

type Cart struct {
	Model
	UserId uint64
	Total  uint16
}

func (Cart) TableName() string {
	return "cart"
}
