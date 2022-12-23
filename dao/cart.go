package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBClient(ctx)}
}

func NewCartTransactionDao(ctx context.Context) *CartDao {
	return &CartDao{NewTransactionDBClient(ctx)}
}

func (dao *CartDao) Create(product *model.Cart) error {
	return dao.DB.Create(&product).Error
}

func (dao *CartDao) GetByUserId(id uint64) (cart *model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("user_id = ?", id).Find(&cart).Error
	return cart, err
}

func (dao *CartDao) UpdateTotal(cart *model.Cart) error {
	return dao.DB.Model(&model.Cart{}).Where("user_id = ?", cart.UserId).Update("total", cart.Total).Error
}
