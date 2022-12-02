package model

import "time"

type MemberAddress struct {
	Id         uint64    `xorm:"pk autoincr comment('用户地址信息Id')"`
	MemberId   uint64    `xorm:"notnull comment('用户Id')"`
	Name       string    `xorm:"varchar(25) notnull comment('收货人姓名')"`
	Tel        string    `xorm:"varchar(11) notnull comment('手机号')"`
	Address    string    `xorm:"varchar(255) notnull comment('地址')"`
	CreateTime time.Time `xorm:"created notnull comment('创建时间')"`
	UpdateTime time.Time `xorm:"updated notnull comment('修改时间')"`
}
