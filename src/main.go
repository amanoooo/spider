package main

import (
	"github.com/amanoooo/spider/src/novel/www/luoxia/com"
	"github.com/amanoooo/spider/src/novel/www/qu/la"
	"github.com/amanoooo/spider/src/photo/miaomi"
	"github.com/amanoooo/spider/src/util"
)

func imageMain() {
	// 猫咪
	miaomi.Main()
}
func novelMain() {
	// 笔趣阁
	go la.Main()
	// 落霞
	com.Main()
}
func main() {
	if util.ENTRY == "photo" {
		imageMain()
	} else {
		novelMain()
	}

	//imageMain()
}
