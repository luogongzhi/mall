package model

import "time"

type Product struct {
	Id            uint64    `xorm:"pk autoincr comment('商品Id')"`
	ProductAttrId uint64    `xorm:"notnull comment('商品属性Id')"`
	Title         string    `xorm:"varchar(255) notnull comment('商品标题')"`
	Info          string    `xorm:"varchar(255) notnull comment('商品描述')"`
	Price         float64   `xorm:"decimal notnull comment('商品价格')"`
	Total         uint64    `xorm:"notnull comment('商品库存')"`
	CreateTime    time.Time `xorm:"created notnull comment('创建时间')"`
	UpdateTime    time.Time `xorm:"updated notnull comment('修改时间')"`
}
