package main

import (
	"errors"
	"fmt"
)

// MySQLError 自定义一个错误结构体类型
type MySQLError struct {
	Code int
	Msg  string
}

// Error 实现error接口要求的方式
func (e *MySQLError) Error() string {
	return e.Msg
}

// 定义包级别的错误
var (
	QueryErr = &MySQLError{Code: 1, Msg: "查询失败"}
	ConnErr  = &MySQLError{Code: 1, Msg: "连接失败"}
)

// 查询数据库的逻辑
func queryDB() error {
	// 自定义错误
	return QueryErr
}

func errorDemo() error {
	// 调用查询数据库的逻辑
	err := queryDB()

	// 基于已有的err错误，包装得到一个新的错误
	newErr := fmt.Errorf("服务端错误，%w", err)
	return newErr
}

func main() {
	err := errorDemo()

	// 被包裹的错误无法使用 == 比较
	if err == QueryErr {
		// ...
	}
	// 剥皮得到里面的err
	innerErr := errors.Unwrap(err)
	fmt.Println(innerErr == QueryErr) // true

	fmt.Println(errors.Is(err, QueryErr)) // true
}
