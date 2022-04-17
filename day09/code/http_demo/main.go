package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// net/http 客户端 发请求 示例

// getDemo get请求示例
func getDemo() {
	resp, err := http.Get("http://liwenzhou.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close() // 关闭

	// resp *http.Response
	// 使用ioutil库读取响应数据
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", b)
}

// getByParam 带参数的GET请求
func getByParam() {
	apiUrl := "http://127.0.0.1:9090/get" // 11.10.6.19/order?orderId=123
	// URL param
	data := url.Values{}
	data.Set("orderId", "123")
	// 自己拼接参数
	// apiUrl = apiUrl + "?" + data.Encode()

	// 使用url包完成参数的拼接
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}
	u.RawQuery = data.Encode() // URL encode
	fmt.Println(u.String())

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("get order failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}

func customClientDemo() {
	header := http.Header{}
	header.Add("key", "value")
	client := &http.Client{
		Transport: &http.Transport{},
	}
	client.Get("http://liwenzhou.com")

	req, _ := http.NewRequest("GET", "http://liwenzhou.com", nil)
	req.Header.Add("key", "value")
}

func postDemo() {
	url := "http://127.0.0.1:9090/post"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=小王子&age=18"

	// json
	contentType := "application/json"
	data := `{"name":"小王子","age":18}` // json格式字符串
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	// 读取响应结果
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}

func urlValuesDemo() {
	var query url.Values
	query.Add("orderId", "123456")
	query.Add("type", "order")

	query.Encode() // http://11.22.33.44/?type=orer&orderId=123456
}

// serverDemo 自己创建一个http server端
func serverDemo() {

	// 做好准备：当客户端请求 127.0.0.1:8082/order?id=1234 会携带一个id
	// 我拿到id去查订单数据给用户返回
	// 注册陆路由方式1
	http.Handle("/order", http.HandlerFunc(f1)) // sum := int64(0)
	// 注册路由方式2
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.ListenAndServe(":8082", nil)
	// http.ListenAndServeTLS()

	// 自定义server
	// s := &http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        http.HandlerFunc(f1),
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()

}

// f1 接收用户的请求，执行业务逻辑，给用户返回响应
func f1(w http.ResponseWriter, r *http.Request) {
	// r:请求
	// r.Body  // 请求体
	// 1.解析请求拿到参数
	id := r.URL.Query().Get("id")
	// 2.业务逻辑处理
	fmt.Println(id)
	// w:响应
	// 3.返回响应
	w.Write([]byte(id))
}

func main() {
	// getDemo()
	// getByParam()

	serverDemo()
}
