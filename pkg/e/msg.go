package e

import "net/http"

var MsgFlags = map[int]string{
	http.StatusOK: "成功",
	InvalidParams: "请求参数错误",
	ErrorDatabase: "数据库操作出错,请重试",

	ErrorExistUser:    "已存在该用户名",
	ErrorNotExistUser: "该用户不存在",
	ErrorNotCompare:   "账号密码错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[http.StatusInternalServerError]
}
