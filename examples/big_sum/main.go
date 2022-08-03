package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_pool_goroutine"
	"sync/atomic"
)

var sum int32 // 总和

// 准备方法
func myFunc(n int32) {
	atomic.AddInt32(&sum, n) // 原子执行相加操作
	fmt.Printf("加上： %d\n", n)
}

func main() {
	// 准备参数
	var args []int32
	for i := 0; i < 1000; i++ {
		args = append(args, int32(i))
	}

	// 批量执行带参数任务
	zdpgo_pool_goroutine.RunBatchArgTask[int32](100, myFunc, args)

	// 查看结果
	fmt.Printf("任务执行完毕，结果是： %d\n", sum)
}
