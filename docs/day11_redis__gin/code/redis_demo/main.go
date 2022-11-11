package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

// redis demo

var (
	rdb *redis.Client // 声明一个全局的redis连接对象
)

// initClient 初始化连接
func initClient() (err error) {
	// 此处应该是初始化全局的redis连接对象
	rdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		// 下面那都是默认值
		Password: "", // no password set
		DB:       0,  // use default DB

		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

// demo1 基本
func demo1() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	// 设置值 类的命令一般用 Err()
	err := rdb.Set(ctx, "name", "wuyong", 10*time.Second).Err()
	fmt.Println(err)

	// 获取值 类的命令后面一般用 Result()
	v, err := rdb.Get(context.Background(), "name").Result()
	if err != nil {
		// 排除掉key不存在的场景
		if err == redis.Nil {
			// 返回的err是key不存在时...
		}
		fmt.Println(err)
		return
	}
	fmt.Println(v, err)

	// 我只想用value,如果出错了就用默认值
	fmt.Println("------")
	fmt.Printf("Err()==redis.Nil:%#v\n", rdb.Get(context.Background(), "namexxxxx").Err() == redis.Nil)
	fmt.Printf("Err()==nil:%#v\n", rdb.Get(context.Background(), "namexxxxx").Err() == nil)
	fmt.Printf("Val():%#v\n", rdb.Get(context.Background(), "namexxxxx").Val())
	nv, nerr := rdb.Get(context.Background(), "namexxxxx").Result()
	fmt.Printf("Result():%#v %#v\n", nv, nerr)
}

func redisExample2() {
	zsetKey := "language:rank"
	languages := []*redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}

	// ZADD
	err := rdb.ZAdd(context.Background(), zsetKey, languages...).Err()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}

	// 把Golang的分数加10
	// err = rdb.ZIncr(context.Background(), zsetKey, &redis.Z{Member: "golang", Score: 1}).Err()
	err = rdb.ZIncrBy(context.Background(), zsetKey, 10.0, "Golang").Err()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}

	// 取分数最高的3个
	ret, err := rdb.ZRevRangeWithScores(context.Background(), zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(context.Background(), zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

func hsetDemo() {
	// 设置hash
	err := rdb.HSet(context.Background(), "wangwenjian", "score", 60).Err()
	fmt.Println(err)

	// 取hash中一个键值对
	v, err := rdb.HGet(context.Background(), "wangwenjian", "weight").Result()
	fmt.Println(v, err)
	// 取多个键值对
	v2, err := rdb.HMGet(context.Background(), "wangwenjian", "age", "weight").Result()
	fmt.Println(v2, err)
}

// pipelineDemo 获取pipeline结果的第一种方式
// 适合命令不多，命令返回的结果格式不太一样
func pipelineDemo() {
	pipe := rdb.Pipeline()
	c1 := pipe.Get(context.Background(), "wangwenjian")
	c2 := pipe.HMGet(context.Background(), "wangwenjian", "age", "weight")
	incr := pipe.Incr(context.Background(), "pipeline_counter")
	pipe.Expire(context.Background(), "pipeline_counter", time.Hour)
	// 执行命令
	_, err := pipe.Exec(context.Background())
	// 取结果
	// 方式1
	fmt.Println(c1.Int())
	fmt.Println(c2.Val())
	fmt.Println(incr.Val(), err)
}

// pipelineDemo2 获取pipeline结果的第二种方式
// 适合批量执行相同类型的命令（返回值类型一致）
func pipelineDemo2() {
	pipe := rdb.Pipeline()
	// 输入命令
	for i := 1; i < 10; i++ {
		pipe.Get(context.Background(), fmt.Sprintf("key%d", i))
	}
	// 执行命令
	cmders, err := pipe.Exec(context.Background())
	if err != nil && err != redis.Nil {
		fmt.Println(err)
		return
	}
	// 取结果
	for _, cmder := range cmders {
		v := cmder.String()
		// 因为拿到的cmder是一个接口类型
		// 需要自己根据上面的命令的返回值进行类型断言
		switch cmder.(type) {
		case *redis.StringCmd:
		case *redis.IntCmd:

			// case *redis.StringSliceCmd:
			// 	// ...

		}
		// cmder.(*redis.StringCmd).Result()
		// cmder.(*redis.StringSliceCmd).Result()
		fmt.Println(v)
	}
	// 方式2

}

func transactionDemo() {
	var (
		maxRetries   = 1000
		routineCount = 10
	)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Increment 使用GET和SET命令以事务方式递增Key的值
	increment := func(key string) error {
		// 事务函数
		txf := func(tx *redis.Tx) error {
			// 获得key的当前值或零值
			n, err := tx.Get(ctx, key).Int()
			if err != nil && err != redis.Nil {
				return err
			}

			// 实际的操作代码（乐观锁定中的本地操作）
			n++

			// 操作仅在 Watch 的 Key 没发生变化的情况下提交
			_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
				pipe.Set(ctx, key, n, 0)
				return nil
			})
			return err
		}

		// 最多重试 maxRetries 次
		for i := 0; i < maxRetries; i++ {
			err := rdb.Watch(ctx, txf, key)
			if err == nil {
				// 成功
				return nil
			}
			if err == redis.TxFailedErr {
				// 乐观锁丢失 重试
				continue
			}
			// 返回其他的错误
			return err
		}

		return errors.New("increment reached maximum number of retries")
	}

	// 模拟 routineCount 个并发同时去修改 counter3 的值
	var wg sync.WaitGroup
	wg.Add(routineCount)
	for i := 0; i < routineCount; i++ {
		go func() {
			defer wg.Done()
			if err := increment("counter33"); err != nil {
				fmt.Println("increment error:", err)
			}
		}()
	}
	wg.Wait()

	n, err := rdb.Get(context.TODO(), "counter33").Int()
	fmt.Println("ended with", n, err)
}

func watchDemo(key string) {
	err := rdb.Watch(context.Background(), func(tx *redis.Tx) error {
		// pipe := tx.TxPipeline()
		// pipe.Set(context.Background(), key, 100, time.Hour)
		// time.Sleep(5 * time.Second) // 假设操作比较耗时
		// _, err := pipe.Exec(context.Background())
		// return err
		_, err := tx.TxPipelined(context.Background(), func(pipe redis.Pipeliner) error {
			pipe.Set(context.Background(), key, 100, time.Hour)
			time.Sleep(5 * time.Second) // 假设操作比较耗时
			_, err := pipe.Exec(context.Background())
			return err
		})
		return err
	}, key)

	fmt.Println(err)
	fmt.Println(rdb.Get(context.Background(), key).Int())
}

func main() {
	// 初始化全局的redis连接对象rdb
	if err := initClient(); err != nil {
		fmt.Println("init redis client failed, err:", err)
		return
	}
	fmt.Println("连接redis成功啦！")

	// 使用全局的redis连接对象执行命令
	// rdb.Set(context.Background(), "name", "杨俊", time.Second)

	// demo1()
	// hsetDemo()

	// transactionDemo()

	watchDemo("lishuo")

}
