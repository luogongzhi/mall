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

func (dao *OrderProductDao) GetListByOrderId(orderId uint64) (returnOrderProduct *[]model.OrderProduct, err error) {
	err = dao.DB.Model(&model.OrderProduct{}).Where("order_id = ?", orderId).Scan(&returnOrderProduct).Error
	return returnOrderProduct, err
}

func (dao *OrderProductDao) DeleteByOrderId(orderId uint64) (err error) {
	return dao.DB.Where("order_id = ?", orderId).Delete(&model.OrderProduct{}).Error
}
