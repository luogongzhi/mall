package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/model"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func NewProductTransactionDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewTransactionDBClient(ctx)}
}

func (dao *ProductDao) GetById(id uint64) (product *model.Product, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Product{}).Where("id = ?", id).Find(&product).Count(&count).Error
	if count == 0 || err != nil {
		return nil, false, err
	}
	return product, true, nil
}

func (dao *ProductDao) Create(product *model.Product) (err error) {
	return dao.DB.Model(&model.Product{}).Create(&product).Error
}

func (dao *ProductDao) UpdateProduct(product *model.Product) (err error) {
	return dao.DB.Select("*").Omit("created_at").Where("id = ?", product.Id).Updates(&product).Error
}

func (dao *ProductDao) UpdateProductTotal(product *model.Product) (err error) {
	return dao.DB.Select("id", "total").Where("id = ?", product.Id).Updates(&product).Error
}

func (dao *ProductDao) Delete(id uint64) (err error) {
	return dao.DB.Where("id = ?", id).Delete(&model.Product{}).Error
}

func (dao *ProductDao) GetList(pageNum, pageSize int) (product []*model.Product, err error) {
	err = dao.DB.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&product).Error
	return product, err
}
