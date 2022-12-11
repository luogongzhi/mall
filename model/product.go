package model

type Product struct {
	Model
	Title     string
	Info      string
	AttrValue string
	Price     float64
	Total     uint64
}

func (Product) TableName() string {
	return "product"
}
