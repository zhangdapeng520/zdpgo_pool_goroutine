package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_pool_goroutine"
	"math/rand"
	"time"
)

func demoFunc() {
	seconds := rand.Intn(30)
	fmt.Println(fmt.Sprintf("模拟一次请求，需要%d秒钟", seconds))
	time.Sleep(time.Second * time.Duration(seconds))
}

func main() {
	// 批量生成任务
	var tasks []func()
	for i := 0; i < 100000; i++ {
		tasks = append(tasks, demoFunc)
	}

	// 批量执行任务
	zdpgo_pool_goroutine.RunBatchTask(tasks)

	// 查看输出
	fmt.Printf("运行中的Gotouine数量: %d\n", zdpgo_pool_goroutine.Running())
	fmt.Printf("任务执行完毕\n")
}
