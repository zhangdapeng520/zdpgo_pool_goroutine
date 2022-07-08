package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_pool_goroutine"
	"math/rand"
	"sync"
)

/*
@Time : 2022/7/8 16:27
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description: 将10000个随机数切分成100份分别计算，然后汇总
*/

const (
	DataSize    = 10000 // 数据个数
	DataPerTask = 100   // 切成多少份
)

// Task 任务对象
type Task struct {
	index int             // 索引
	nums  []int           // 数据
	sum   int             // 总共
	wg    *sync.WaitGroup // 等待组
}

// Do 执行任务
func (t *Task) Do() {
	// 求和
	for _, num := range t.nums {
		t.sum += num
	}

	// 等待组完成
	t.wg.Done()
}

// 任务方法
func taskFunc(data interface{}) {
	// 转换为任务对象
	task := data.(*Task)

	// 执行任务
	task.Do()

	// 查看结果
	fmt.Printf("task:%d sum:%d\n", task.index, task.sum)
}

func main() {
	// 创建协程池
	p, _ := zdpgo_pool_goroutine.NewPoolWithFunc(10, taskFunc)
	defer p.Release()

	// 创建随机数
	nums := make([]int, DataSize, DataSize)
	for i := range nums {
		nums[i] = rand.Intn(1000)
	}

	// 创建等待组，并添加指定个数的任务
	var wg sync.WaitGroup
	wg.Add(DataSize / DataPerTask)

	// 创建任务组
	tasks := make([]*Task, 0, DataSize/DataPerTask)

	// 实例化任务
	for i := 0; i < DataSize/DataPerTask; i++ {
		task := &Task{
			index: i + 1,
			nums:  nums[i*DataPerTask : (i+1)*DataPerTask],
			wg:    &wg,
		}
		tasks = append(tasks, task)
		p.Invoke(task) // 执行任务
	}

	// 等待所有任务执行完毕
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", zdpgo_pool_goroutine.Running())

	// 统计结果
	var sum int
	for _, task := range tasks {
		sum += task.sum
	}

	// 期望值
	var expect int
	for _, num := range nums {
		expect += num
	}

	// 比较结果
	fmt.Printf("finish all tasks, result is %d expect:%d\n", sum, expect)
}
