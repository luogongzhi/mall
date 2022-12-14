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

func (dao *CartDao) CreateCart(product *model.Cart) error {
	return dao.DB.Create(&product).Error
}
