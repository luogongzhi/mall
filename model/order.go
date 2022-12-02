package model

import "time"

type Order struct {
	Id            uint64    `xorm:"pk autoincr comment('订单Id')"`
	MemberId      uint64    `xorm:"notnull comment('用户Id')"`
	AddressId     uint64    `xorm:"notnull comment('收货地址信息Id')"`
	ProductId     string    `xorm:"longtext notnull comment('分号分隔的商品Id')"`
	ProductAmount float64   `xorm:"decimal notnull comment('商品总额')"`
	FreightAmount float64   `xorm:"decimal notnull comment('运费金额')"`
	TotalAmount   float64   `xorm:"decimal notnull comment('订单总额')"`
	Status        bool      `xorm:"tinyint(2) notnull comment('订单状态 0->未完成 1->已完成')"`
	CreateTime    time.Time `xorm:"created notnull comment('创建时间')"`
	UpdateTime    time.Time `xorm:"updated notnull comment('修改时间')"`
}
