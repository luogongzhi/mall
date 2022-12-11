package model

type UserAddress struct {
	Model
	UserId         uint64
	Name           string
	Tel            string
	AddressDetails string
}

func (UserAddress) TableName() string {
	return "user_address"
}
