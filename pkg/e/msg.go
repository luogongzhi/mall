package e

import "net/http"

var MsgFlags = map[int]string{
	InvalidParams: "参数有误",
	http.StatusOK: "成功",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[http.StatusInternalServerError]
}
