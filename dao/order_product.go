package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
)

type OrderProductDao struct {
	*gorm.DB
}

func NewOrderProductDao(ctx context.Context) *OrderProductDao {
	return &OrderProductDao{NewDBClient(ctx)}
}

func (dao *OrderProductDao) Create(orderProduct *model.OrderProduct) (err error) {
	return dao.DB.Model(&model.OrderProduct{}).Create(&orderProduct).Error
}
