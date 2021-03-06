package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_pool_goroutine"
	"math/rand"
	"sync"
	"time"
)

func demoFunc() {
	seconds := rand.Intn(30)
	fmt.Println(fmt.Sprintf("模拟一次请求，需要%d秒钟", seconds))
	time.Sleep(time.Second * time.Duration(seconds))
}

func main() {
	// 释放ants的默认协程池
	defer zdpgo_pool_goroutine.Release()
	var wg sync.WaitGroup

	// 任务函数
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		_ = zdpgo_pool_goroutine.Submit(syncCalculateSum) // 提交任务到默认协程池
	}
	wg.Wait()

	// 查看输出
	fmt.Printf("运行中的Gotouine数量: %d\n", zdpgo_pool_goroutine.Running())
	fmt.Printf("任务执行完毕\n")
}
