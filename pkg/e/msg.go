package e

import "net/http"

var MsgFlags = map[int]string{
	http.StatusOK: "成功",
	InvalidParams: "请求参数错误",
	ErrorDatabase: "数据库操作出错,请重试",
	ErrorDate:     "日期错误",

	ErrorExistUser:           "已存在该用户名",
	ErrorNotExistUser:        "该用户不存在",
	ErrorNotCompare:          "账号密码错误",
	ErrorNotExistUserAddress: "该用户地址不存在",

	ErrorAuthCheckTokenFail:        "Token鉴权失败",
	ErrorAuthCheckTokenTimeout:     "Token已超时",
	ErrorAuthToken:                 "Token生成失败",
	ErrorAuth:                      "Token错误",
	ErrorAuthInsufficientAuthority: "权限不足",

	ErrorNotExistProduct:  "该商品不存在",
	ErrorProductTotal:     "错误的商品数量",
	ErrorNotEnoughProduct: "该商品库存不足",

	ErrorNotExistCartProduct: "购物车中未添加商品",

	ErrorNotExistOrder: "订单不存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[http.StatusInternalServerError]
}
