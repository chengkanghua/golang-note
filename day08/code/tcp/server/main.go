package main

import (
	"bufio"
	"fmt"
	"net"
)

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		// 基于网络连接创建一个reader对象
		reader := bufio.NewReader(conn)
		// 每次读取128字节
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		// 把收到的内容转为字符串
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		// 把收到的内容发送给客户端
		conn.Write([]byte(recvStr)) // 发送数据
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20001")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}
