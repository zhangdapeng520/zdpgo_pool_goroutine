package zdpgo_pool_goroutine

import (
	"sync"
)

// RunBatchTask 批量执行任务
func RunBatchTask(funcList []func()) {
	defer Release() // 释放ants的默认协程池
	var wg sync.WaitGroup

	for _, funcObj := range funcList {
		wg.Add(1)

		taskFunc := func() {
			funcObj()
			wg.Done()
		}

		_ = Submit(taskFunc) // 提交任务到默认协程池
	}

	wg.Wait()
}

// RunBatchArgTask 批量执行带参数的任务
// @param poolSize 协程池的容量大小
// @param funcObj 要执行的带参方法
// @param args 参数列表，会将每个参数都传给funcObj并发执行
func RunBatchArgTask[T comparable](poolSize int, funcObj func(arg T), args []T) {
	var wg sync.WaitGroup

	// 初始化协程池
	p, _ := NewPoolWithFunc(poolSize, func(i interface{}) {
		funcObj(i.(T)) // Invoke的参数是interface{}类型，这里需要重新转换为T类型
		wg.Done()
	})

	// 释放协程池
	defer p.Release()

	// 提交任务
	for _, arg := range args {
		wg.Add(1)
		_ = p.Invoke(arg)
	}

	wg.Wait()
}
