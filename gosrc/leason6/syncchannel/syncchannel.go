package main

import (
	"fmt"
	"sync"
)

type FuncHandler struct {
	ch      chan struct{}
	Handler func() string
}

var handlerMap map[int]*FuncHandler

var wg sync.WaitGroup
var lock sync.Mutex

// 这个里面的waitgroup是为了让主线程不退出等待子协程执行结束以后再退出，是一种同步机制
// dofunc的意义在于，让每次function调用都可以获取一个i的程序拷贝，goroutine的创建速度满与for循环的执行，会出现所有goroutine都将i看成9的风险，goroutine是一种异步机制
// mutex是一个互斥锁，是一种互斥机制
// chan在这里起到的是一种同步机制
func main() {
	handlerMap = make(map[int]*FuncHandler)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		dofunc(i)
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		dofunc(i)
	}
	wg.Wait()
}

func dofunc(i int) {
	go func() {
		lock.Lock()
		if val, ok := handlerMap[i]; !ok {
			fh := &FuncHandler{}
			fh.ch = make(chan struct{})
			handlerMap[i] = fh
			lock.Unlock()
			fh.Handler = func() string {
				// 在这里你可以做复杂的异步流程
				return fmt.Sprintf("i am %d", i)
			}
			fmt.Printf("input: i am %d\n", i)
			close(fh.ch)
		} else {
			lock.Unlock()
			<-val.ch // kv 对的锁
			fmt.Println("output: ", val.Handler())
		}
		wg.Done()
	}()
}
