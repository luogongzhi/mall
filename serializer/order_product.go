package serializer

type OrderProductVO struct {
	ProductId uint64  `json:"product_id"`
	Title     string  `json:"title"`
	Info      string  `json:"info"`
	AttrValue string  `json:"attr_value"`
	Price     float64 `json:"price"`
	Total     uint16  `json:"total"`
}
