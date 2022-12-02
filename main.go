package main

import (
	_ "mall/dao"
	"mall/route"
)

func main() {
	r := route.Router()
	r.Run()
}
