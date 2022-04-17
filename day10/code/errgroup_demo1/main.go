package main

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/sync/errgroup"
)

func f1() error {
	return errors.New("f1 error")
}

// fetchUrlDemo 并发获取url内容
func fetchUrlDemo() {
	wg := sync.WaitGroup{}
	var urls = []string{
		"http://pkg.go.dev",
		"http://www.liwenzhou.com",
		"http://www.yixieqitawangzhi.com",
	}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err == nil {
				fmt.Printf("获取%s成功\n", url)
				resp.Body.Close()
			}
			return // 如何将错误返回呢？
		}(url)
	}
	wg.Wait()
	// 如何获取goroutine中可能出现的错误呢？
}

func main() {
	// go err := f1()
	// fmt.Println(err)

	// var ch = make(chan error)
	// go func() {
	// 	err := f1()
	// 	ch <- err
	// }()

	// for err := range ch {
	// 	fmt.Println(err)
	// }

	// fetchUrlDemo2()
	f3()
}

// fetchUrlDemo2 使用errgroup并发获取url内容
func fetchUrlDemo2() error {
	g := new(errgroup.Group) // 创建等待组（类似sync.WaitGroup）

	var urls = []string{
		"http://pkg.go.dev",
		"http://www.liwenzhouxx.com",
		"http://www.yixieqitawangzhi.com",
	}
	// 现代版刻舟求剑
	for _, url := range urls {
		url := url // 注意此处声明新的变量
		// 启动一个goroutine去获取url内容
		g.Go(func() error {
			// fmt.Printf("%p\n", &url)
			resp, err := http.Get(url)
			if err == nil {
				fmt.Printf("获取%s成功\n", url)
				resp.Body.Close()
			}
			return err // 返回错误
		})
	}
	// Wait会等待所有的goroutine结束
	if err := g.Wait(); err != nil {
		// 只返回第一个错误
		fmt.Println(err)
		return err
	}
	fmt.Println("所有goroutine均成功")
	return nil
}

func f3() {
	var urls = []string{
		"http://pkg.go.dev",
		"http://www.liwenzhouxx.com",
		"http://www.yixieqitawangzhi.com",
	}
	var wg sync.WaitGroup
	for _, url := range urls {
		url := url // 每次for循环生成了一个新的局部变量
		wg.Add(1)
		go func() {
			fmt.Printf("%p\n", &url)
			wg.Done()
		}()
	}
	wg.Wait()
}
