package model

type MemberAddress struct {
	Model
	MemberId       uint64
	Name           string
	Tel            string
	AddressDetails string
}

func (MemberAddress) TableName() string {
	return "member_address"
}
