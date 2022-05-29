package main

import (
	"flag"
	"fmt"
	"gin_viper_mysql/service"
	"reflect"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() (err error) {
	// dsn := "root:123@tcp(10.211.55.6:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	//         root:123@tcp(10.221.55.6:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local
	// viper 配置文件
	viper.SetConfigFile("./config/config.yaml")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		// panic(err)
		return
	}
	fmt.Println("--->", viper.GetString("mysql.password"))

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%v",
		viper.Get("mysql.user"),
		viper.GetString("mysql.password"),
		viper.Get("mysql.host"),
		viper.Get("mysql.port"),
		viper.Get("mysql.database"),
		viper.Get("mysql.config"),
	)
	fmt.Println(dsn)                 //root:123@tcp(10.221.55.6:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local
	fmt.Println(reflect.TypeOf(dsn)) //String

	// 初始化全局的db
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}

func main() {
	var name string
	flag.StringVar(&name, "name", "杨俊", "请输入名字")
	flag.Parse()
	service.InitName(name)

	fmt.Println(name)

	// 连接数据库
	// if err := initDB(); err != nil {
	// 	fmt.Println("connect mysql failed", err)
	// 	panic(err)
	// } else {
	// 	fmt.Println("connect mysql success")
	// }

}
