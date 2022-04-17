package main

import (
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/sync/errgroup"
)

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

func fetchUrlDemo2() error {
	g := new(errgroup.Group) // 创建等待组（类似sync.WaitGroup）
	var urls = []string{
		"http://pkg.go.dev",
		"http://www.liwenzhou.com",
		"http://www.yixieqitawangzhi.com",
	}
	for _, url := range urls {
		url := url
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				fmt.Printf("获取%s成功\n", url)
				resp.Body.Close()
			}
			return err
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("所有goroutine均成功")
	return nil
}

func main() {
	// fetchUrlDemo()
	fetchUrlDemo2()
}
