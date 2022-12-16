package serializer

type OrderProductVO struct {
	ProductId uint64 `json:"product_id"`
	Total     uint16 `json:"total"`
}
