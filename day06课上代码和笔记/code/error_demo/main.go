package main

import (
	"errors"
	"fmt"
	"os"
)

// error

var (
	// 使用errors.New 创建error 包含一段错误描述信息
	ErrInvalidOp = errors.New("无效的操作")

	ErrInvalidId = errors.New("无效的订单号")
)

type Order struct{}

// queryOrder 根据订单id查询订单数据的函数
func queryOrder(id int64) (*Order, error) {
	// 去数据库查询
	if err := ConnectDB("jade", "111"); err != nil {
		// return nil, errors.New("服务端异常")
		// return nil, errors.New(fmt.Sprintf("连接数据库失败，err:%v", err))
		// return nil, fmt.Errorf("连接数据库失败，err:%v", err)
		// 既能拿到一个新的error,又不丢失原来的error
		// 基于一个已有的从 error 包装得到一个新的 error

		// 使用 fmt.Errorf 搭配特殊的格式化动词 %w
		return nil, fmt.Errorf("连接数据库失败，err:%w", err)
	}

	// 连接数据库正常，但是没查到这个订单号对应的订单信息
	if false {
		return nil, ErrInvalidId
	}

	// 查询成功
	return &Order{}, nil
}

func main() {
	// 调用标准库的Open函数打开文件
	// f, err := os.Open("xx.txt")
	// if err != nil { // error 接口类型默认零值为nil
	// 	// 说明有错误
	// 	fmt.Println(err) // 会调用 err.Error() 拿到 错误描述信息
	// 	return
	// }
	// // 代码能执行到这里说明打开文件成功，f是个正经文件对象
	// defer f.Close()

	// ErrInvalidOp.Error()

	order, err := queryOrder(10086)
	if err != nil {
		fmt.Println(err == ErrConnFailed)

		oErr := errors.Unwrap(err) // 解包
		fmt.Println(oErr, oErr == ErrConnFailed)

		if ok := errors.Is(err, ErrConnFailed); ok { // 是否包含
			fmt.Println("是底层连接数据库失败了")
		}

		var nErr *DBError
		if ok := errors.As(err, &nErr); ok {
			// 自动做error的类型转换
		}

		fmt.Println(err)
		if err == ErrInvalidId {
			// 提示用户订单号输入有误
			// log.Warning(err)
		}
		// 查询失败
	}
	fmt.Println(order)

	os.OpenFile("./main.go", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
}
