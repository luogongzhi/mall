package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
	"mall/serializer"
)

type CartProductDao struct {
	*gorm.DB
}

func NewCartProductDao(ctx context.Context) *CartProductDao {
	return &CartProductDao{NewDBClient(ctx)}
}

func (dao *CartProductDao) CreateCartProduct(cartProduct *model.CartProduct) error {
	return dao.DB.Create(&cartProduct).Error
}

func (dao *CartProductDao) GetCartProductList(cartId uint64) (cartProductVOList *[]serializer.CartProductVO, err error) {
	err = dao.DB.Model(&model.CartProduct{}).Select("cart_product.product_id, product.title, product.info, product.attr_value, product.price, cart_product.total").
		Joins("left join product on cart_product.product_id = product.id").Where("cart_product.cart_id = ?", cartId).
		Scan(&cartProductVOList).Error
	return cartProductVOList, err
}
