package utils

import "mall/dao"

// Rollback 业务回滚
func Rollback(args ...interface{}) {
	for _, v := range args {
		switch v.(type) {
		case *dao.UserDao:
			v.(*dao.UserDao).Rollback()
		case *dao.CartDao:
			v.(*dao.CartDao).Rollback()
		case *dao.CartProductDao:
			v.(*dao.CartProductDao).Rollback()
		case *dao.ProductDao:
			v.(*dao.ProductDao).Rollback()
		case *dao.OrderDao:
			v.(*dao.OrderDao).Rollback()
		case *dao.OrderProductDao:
			v.(*dao.OrderProductDao).Rollback()
		}
	}
}

// Commit 业务提交
func Commit(args ...interface{}) {
	for _, v := range args {
		switch v.(type) {
		case *dao.UserDao:
			v.(*dao.UserDao).Commit()
		case *dao.CartDao:
			v.(*dao.CartDao).Commit()
		case *dao.CartProductDao:
			v.(*dao.CartProductDao).Commit()
		case *dao.ProductDao:
			v.(*dao.ProductDao).Commit()
		case *dao.OrderDao:
			v.(*dao.OrderDao).Commit()
		case *dao.OrderProductDao:
			v.(*dao.OrderProductDao).Commit()
		}
	}
}
