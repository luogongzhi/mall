package model

import "time"

type ProductAttr struct {
	Id         uint64    `xorm:"pk autoincr comment('商品属性Id')"`
	AttrValue  string    `xorm:"varchar(255) notnull comment('属性值')"`
	CreateTime time.Time `xorm:"created notnull comment('创建时间')"`
	UpdateTime time.Time `xorm:"updated notnull comment('修改时间')"`
}
