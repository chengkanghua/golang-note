package main

import (
	"fmt"
	"log"

	"github.com/philchia/agollo/v4"
)

func main() {
	// 连接apollo
	agollo.Start(&agollo.Conf{
		AppID:           "SampleApp",
		Cluster:         "dev",
		NameSpaceNames:  []string{"application.properties", "shopping_cart.yaml"},
		MetaAddr:        "http://localhost:8080",
		AccesskeySecret: "b8ceb3ec62f34030b1b1fd9a431e420b",
	})
	// 监听配置修改
	agollo.OnUpdate(func(event *agollo.ChangeEvent) {
		// 监听配置变更
		log.Printf("event:%#v\n", event)
	})

	log.Println("初始化Apollo配置成功")

	// 从默认的application.properties命名空间获取key的值
	val := agollo.GetString("timeout")
	log.Println(val)
	// 获取命名空间下所有key
	keys := agollo.GetAllKeys(agollo.WithNamespace("shopping_cart.yaml"))
	fmt.Println(keys)
	// 获取指定一个命令空间下key的值
	other := agollo.GetString("content", agollo.WithNamespace("shopping_cart.yaml"))
	log.Println(other)
	// 获取指定命名空间下的所有内容
	namespaceContent := agollo.GetContent(agollo.WithNamespace("shopping_cart.yaml"))
	log.Println(namespaceContent)
	select {}
}
