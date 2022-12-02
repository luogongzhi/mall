package model

import "time"

type Member struct {
	Id         uint64    `xorm:"pk autoincr comment('用户Id')"`
	Username   string    `xorm:"varchar(25) notnull unique comment('用户名')"`
	Password   string    `xorm:"varchar(32) notnull comment('密码')"`
	Tel        string    `xorm:"varchar(11) comment('手机号')"`
	Email      string    `xorm:"notnull varchar(20) comment('邮箱地址')"`
	Gender     bool      `xorm:"tinyint(3) default(0) comment('性别 0->未知 1->男 2->女')"`
	Birth      time.Time `xorm:"notnull comment('生日')"`
	CreateTime time.Time `xorm:"created notnull comment('创建时间')"`
	UpdateTime time.Time `xorm:"updated notnull comment('修改时间')"`
}
