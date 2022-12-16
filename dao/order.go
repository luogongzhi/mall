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

func (dao *OrderDao) Create(order *model.Order) (id uint64, err error) {
	err = dao.DB.Model(&model.Order{}).Create(&order).Error
	return order.Model.Id, err
}

func (dao *OrderDao) GetByOrderId(userId, orderId uint64) (order *model.Order, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Order{}).Where("user_id = ? AND id = ?", userId, orderId).Find(&order).Count(&count).Error
	if count == 0 || err != nil {
		return nil, false, err
	}
	return order, true, nil
}

func (dao *OrderDao) Update(order *model.Order) error {
	return dao.DB.Where("id = ?", order.Id).Updates(&order).Error
}

func (dao *OrderDao) GetListByUserId(userId uint64) (orderList *[]model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("user_id = ?", userId).Scan(&orderList).Error
	return orderList, err
}

func (dao *OrderDao) DeleteById(id uint64) (err error) {
	return dao.DB.Where("id = ?", id).Delete(&model.Order{}).Error
}
