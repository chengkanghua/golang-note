package blog

import (
	"fmt"
	"gomod_demo/app/shop"
)

func Demo() {
	fmt.Println("app/blog/blog.go")
}

// CallShopDemo 调用其他包中的函数
func CallShopDemo() {
	shop.Demo()
}
