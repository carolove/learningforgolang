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
			<-val.ch
			fmt.Println("output: ", val.Handler())
		}
		wg.Done()
	}()
}
