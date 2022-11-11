package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Conf struct {
	Version string               `mapstructure:"version"`
	Port    int                  `mapstructure:"port"`
	App     `mapstructure:"app"` // 嵌套
}

type App struct {
	Node int `mapstructure:"node"`
}

func main() {
	// 1.指定配置文件来源
	viper.SetConfigFile("./config.yaml")
	// viper.SetConfigName("config")
	// viper.AddConfigPath("./conf")
	// viper.SetConfigType("yaml")
	// ...

	// 2.从配置文件读取配置信息
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	viper.WatchConfig() // 持续监视着配置文件
	// 还可以在配置文件发生变化的时候 执行回调函数（自定义操作）
	viper.OnConfigChange(callbackFunc) // 把我写的回调函数告诉viper

	// 每隔3秒钟执行
	// for t := range time.Tick(3 * time.Second) {
	// 	fmt.Printf(
	// 		"%s:%v\n",
	// 		t.Format("2006/01/02 15:04:05"),
	// 		viper.GetString("version"),
	// 	)
	// }

	// 从环境变量加载配置
	viper.SetEnvPrefix("todo") // => TODO_
	viper.BindEnv("name")
	fmt.Println(viper.Get("name")) // 从环境变量里取配置  => TODO_NAME

	// 从命令行获取参数
	// 1. 登记那些参数是从命令行获取
	pflag.Int("xxx", 1234, "帮助信息")

	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	i := viper.GetInt("xxx") // 从viper而不是从pflag检索值
	fmt.Println("从命令行获取到xxx:", i)

	// 定义一个结构体变量
	var cfg Conf
	// 把viper读取的配置反序列化到cfg变量里
	err = viper.Unmarshal(&cfg)
	fmt.Println(err)
	fmt.Printf("cfg:%#v err:%v\n", cfg, err)

	// 3.使用配置
	fmt.Println(viper.Get("port").(int)) // Get返回的是空接口类型
	fmt.Println(viper.GetInt("port"))    // GetInt直接返回int类型（默认0）
	fmt.Println(viper.Get("version"))
	fmt.Println(viper.GetString("version"))
	fmt.Println(viper.GetInt("app.node"))

	r := gin.Default()

	// gin框架启动时候的端口号
	r.Run(fmt.Sprintf(":%d", viper.GetInt("port")))

}

// callbackFunc 当配置文件发生变化的时候，viper执行的回调函数
func callbackFunc(e fsnotify.Event) {
	// 配置文件发生变更之后会调用的回调函数
	fmt.Printf("Config file changed:%#v\n", e)
	if e.Op == fsnotify.Write {
		// ....
	}
}
