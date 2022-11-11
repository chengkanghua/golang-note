package main

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// gorm demo

// 定义一个结构体对应的数据库里的一张表
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	gorm.Model
	Name string
	Age  *int `gorm:"default:18"`
	// Age    sql.NullInt64 `gorm:"default:18"`
	Active sql.NullBool `gorm:"default:true"`
}

func main() {
	// "" 0 false
	var p11 = User{Name: "p11"} // p11.Active=false
	var age int = 20
	var p12 = User{Name: "p12", Age: &age}
	// var p12 = User{Name: "p12", Active: sql.NullBool{Bool: false, Valid: true}}
	fmt.Println(p11, p12)

	// 连接数据库
	dsn := "root:root1234@tcp(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("gorm.Open failed", err)
		panic(err)
	}

	db.AutoMigrate(&Product{})

	// 增删改查
	// p1 := &Product{
	// 	Code:  "D42",
	// 	Price: 100,
	// }
	// db.Create(p1)
	// 查询单条记录

	var p2 Product
	db.First(&p2, 1) // select * from products where id=1;
	fmt.Printf("p2:%#v\n", p2)

	db.First(&p2, "code = ?", "D42") // select * from products where code="D42";
	fmt.Printf("p2:%#v\n", p2)

	// 更新
	// Update - 将 p2 的 price 更新为 200
	db.Debug().Model(&p2).Update("Price", 200)
	db.Model(&p2).Updates(map[string]interface{}{
		"Code":  "X42",
		"Price": 250,
	})

}
