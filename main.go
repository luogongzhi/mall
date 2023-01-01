package main

import (
	"mall/conf"
	"mall/route"
)

func main() {
	conf.Init()
	r := route.Router()
	_ = r.Run()
}
