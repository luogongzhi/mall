package model

import "time"

type Cart struct {
	Id         uint64    `xorm:"pk autoincr comment('购物车Id')"`
	MemberId   uint64    `xorm:"notnull comment('用户Id')"`
	ProductId  string    `xorm:"longtext notnull comment('分号分隔的商品Id')"`
	Total      uint16    `xorm:"int(10) notnull comment('商品数量')"`
	CreateTime time.Time `xorm:"created notnull comment('创建时间')"`
	UpdateTime time.Time `xorm:"updated notnull comment('修改时间')"`
}
