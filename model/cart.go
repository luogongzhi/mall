package model

type Cart struct {
	Model
	MemberId uint64
	Total    uint16
}

func (Cart) TableName() string {
	return "cart"
}
