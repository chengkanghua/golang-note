package main

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// gorm 示例

type Book struct {
	gorm.Model
	Title  string  `gorm:"title"`
	Amount int     `gorm:"amount"`
	Price  float64 `gorm:"price"`
	Status bool    `gorm:"status"`
}

// updateDemo gorm更新示例
func updateDemo() {
	// 全都改了
	db.Debug().
		Model(&Book{}).
		Updates(Book{Title: "hello", Amount: 18, Status: false})

	// 1. 当你正好有一个结构体对象（包含数据库主键） -> 能够对应到数据库里的一条记录
	// 接口幂等 -> 先查（是否存在这条记录；做状态判断）再更新
	var id uint = 1
	var b1 Book
	err := db.Where("id = ?", id).First(&b1).Error // 先查一条记录（包含主键）
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("参数错误")
	}

	db.Debug().
		Model(&b1).
		Updates(Book{Title: "hello2", Amount: 28, Status: false})

		// 2. 直接通过where条件来查找记录并更新
	cond := &Book{
		Model: gorm.Model{ // 匿名嵌入的结构体，字段名默认就是类型名
			ID: id,
		},
	}
	db.Debug().
		Where(cond).
		Updates(Book{Title: "hello3", Amount: 38, Status: false})

}

// initDB 初始化数据库
func initDB(dsn string) (err error) {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}

func main() {
	dsn := "root:root1234@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	if err := initDB(dsn); err != nil {
		panic(err)
	}

	db.AutoMigrate(&Book{})

	// b1 := Book{
	// 	Title:  "《跟七米学Go语言2》",
	// 	Amount: 100,
	// 	Price:  19.99,
	// 	Status: false,
	// }

	// db.Create(&b1)

	// 数据库操作
	updateDemo()
}
