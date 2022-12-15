package serializer

import "mall/model"

type ProductVO struct {
	Id        uint64  `json:"id"`
	Title     string  `json:"title"`
	Info      string  `json:"info"`
	Price     float64 `json:"price"`
	Total     uint16  `json:"total"`
	AttrValue string  `json:"attr_value"`
}

type ProductCreateDTO struct {
	Title     string  `json:"title" binding:"required"`
	Info      string  `json:"info" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
	Total     uint16  `json:"total" binding:"required"`
	AttrValue string  `json:"attr_value" binding:"required"`
}

type ProductUpdateDTO struct {
	Id        uint64  `json:"id" binding:"required"`
	Title     string  `json:"title" binding:"required"`
	Info      string  `json:"info" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
	Total     uint16  `json:"total" binding:"required"`
	AttrValue string  `json:"attr_value" binding:"required"`
}

type ProductDeleteDTO struct {
	Id uint64 `json:"id" binding:"required"`
}

func NewProductVO(product *model.Product) ProductVO {
	return ProductVO{
		Id:        product.Id,
		Title:     product.Title,
		Info:      product.Info,
		Price:     product.Price,
		Total:     product.Total,
		AttrValue: product.AttrValue,
	}
}

func NewProductVOList(productList []*model.Product) (productVOList []ProductVO) {
	for _, item := range productList {
		product := NewProductVO(item)
		productVOList = append(productVOList, product)
	}
	return productVOList
}
