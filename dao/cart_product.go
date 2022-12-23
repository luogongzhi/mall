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

func NewCartProductTransactionDao(ctx context.Context) *CartProductDao {
	return &CartProductDao{NewTransactionDBClient(ctx)}
}

func (dao *CartProductDao) Create(cartProduct *model.CartProduct) error {
	return dao.DB.Create(&cartProduct).Error
}

func (dao *CartProductDao) GetList(cartId uint64) (cartProductVOList *[]serializer.CartProductVO, err error) {
	err = dao.DB.Model(&model.CartProduct{}).Select("cart_product.product_id, product.title, product.info, product.attr_value, product.price, cart_product.total").
		Joins("left join product on cart_product.product_id = product.id").Where("cart_product.cart_id = ?", cartId).
		Scan(&cartProductVOList).Error
	return cartProductVOList, err
}

func (dao *CartProductDao) UpdateTotalByProductId(cartProduct *model.CartProduct) error {
	return dao.DB.Where("cart_id = ? AND product_id = ?", cartProduct.CartId, cartProduct.ProductId).Updates(&cartProduct).Error
}

func (dao *CartProductDao) DeleteByProductId(cartProduct *model.CartProduct) error {
	return dao.DB.Where("cart_id = ? AND product_id = ?", cartProduct.CartId, cartProduct.ProductId).Delete(&model.CartProduct{}).Error
}

func (dao *CartProductDao) DeleteByCartId(cartId uint64) error {
	return dao.DB.Where("cart_id = ?", cartId).Delete(&model.CartProduct{}).Error
}
func (dao *CartProductDao) GetByProductId(cartProduct *model.CartProduct) (returnCartProduct *model.CartProduct, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.CartProduct{}).Where("cart_id = ? AND product_id = ?", cartProduct.CartId, cartProduct.ProductId).Find(&returnCartProduct).Count(&count).Error
	if count == 0 || err != nil {
		return nil, false, err
	}
	return returnCartProduct, true, nil
}
