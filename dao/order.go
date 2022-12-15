package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

func (dao *OrderDao) Create(orderProduct *model.Order) (id uint64, err error) {
	err = dao.DB.Model(&model.Order{}).Create(&orderProduct).Error
	return orderProduct.Model.Id, err
}
