package main

import (
	"mall/conf"
	_ "mall/dao"
	"mall/route"
)

func main() {
	conf.Init()
	r := route.Router()
	r.Run()
}
