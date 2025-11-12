package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	test2()
}

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func test1() {
	mutex := sync.Mutex{}

	count := 0
	wg := sync.WaitGroup{}
	for j := 0; j < 10; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			for i := 0; i < 1000; i++ {
				count++
			}
		}()
	}
	wg.Wait()

	fmt.Println("count:", count)
}

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func test2() {
	count := atomic.Int32{}

	wg := sync.WaitGroup{}
	for j := 0; j < 10; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				count.Add(1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("count:", count.Load())
}
