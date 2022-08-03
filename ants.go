package zdpgo_pool_goroutine

import (
	"errors"
	"log"
	"os"
	"runtime"
	"time"
)

const (
	defaultGoroutinePoolSize = 333333          // 默认的协程池数量
	DefaultCleanIntervalTime = time.Second * 3 // 默认的清除间隔时间
	OPENED                   = iota            // 连接池开启状态
	CLOSED                                     // 连接池关闭状态
)

var (
	// ErrLackPoolFunc will be returned when invokers don't provide function for pool.
	ErrLackPoolFunc = errors.New("must provide function for pool")

	// ErrInvalidPoolExpiry will be returned when setting a negative number as the periodic duration to purge goroutines.
	ErrInvalidPoolExpiry = errors.New("invalid expiry for pool")

	// ErrPoolClosed will be returned when submitting task to a closed pool.
	ErrPoolClosed = errors.New("this pool has been closed")

	// ErrPoolOverload will be returned when the pool is full and no workers available.
	ErrPoolOverload = errors.New("too many goroutines blocked on submit or Nonblocking is set")

	// ErrInvalidPreAllocSize will be returned when trying to set up a negative capacity under PreAlloc mode.
	ErrInvalidPreAllocSize = errors.New("can not set up a negative capacity under PreAlloc mode")

	// ErrTimeout will be returned after the operations timed out.
	ErrTimeout = errors.New("operation timed out")

	// workerChanCap determines whether the channel of a worker should be a buffered channel
	// to get the best performance. Inspired by fasthttp at
	// https://github.com/valyala/fasthttp/blob/master/workerpool.go#L139
	workerChanCap = func() int {
		// Use blocking channel if GOMAXPROCS=1.
		// This switches context from sender to receiver immediately,
		// which results in higher performance (under go1.5 at least).
		if runtime.GOMAXPROCS(0) == 1 {
			return 0
		}

		// Use non-blocking workerChan if GOMAXPROCS>1,
		// since otherwise the sender might be dragged down if the receiver is CPU-bound.
		return 1
	}()

	defaultLogger = Logger(log.New(os.Stderr, "", log.LstdFlags))

	// 当此库被导入的时候，会初始化一个Goroutine协程池
	defaultGoroutinePool, _ = NewPool(defaultGoroutinePoolSize)
)

// Logger 用于日志格式化
type Logger interface {
	// Printf 必须和 log.Printf 具有相同的实现
	Printf(format string, args ...interface{})
}

// Submit 提交任务到连接池
func Submit(task func()) error {
	return defaultGoroutinePool.Submit(task)
}

// Running 返回当前整型运行的Goroutine的数量
func Running() int {
	return defaultGoroutinePool.Running()
}

// Cap 返回默认连接池的容量
func Cap() int {
	return defaultGoroutinePool.Cap()
}

// Free 获取可用的Goroutine数量
func Free() int {
	return defaultGoroutinePool.Free()
}

// Release 关闭默认的连接池
func Release() {
	defaultGoroutinePool.Release()
}

// Reboot 重启默认的连接池
func Reboot() {
	defaultGoroutinePool.Reboot()
}
