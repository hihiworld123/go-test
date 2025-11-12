package main

import (
	"fmt"
	"time"
)

func main() {
	test2()
}

// 编写一个程序，使用通道实现两个协程之间的通信。
// 一个协程生成从1到10的整数，并将这些整数发送到通道中，
// 另一个协程从通道中接收这些整数并打印出来。
func test1() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			select {
			case i := <-ch:
				fmt.Printf("%d ", i)
			case <-time.After(1 * time.Second):
				fmt.Println("timeout")
			}
		}

	}()

	time.Sleep(1 * time.Second)
}

// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，
// 消费者协程从通道中接收这些整数并打印。
func test2() {
	ch := make(chan int, 1)

	//生产者
	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
	}()

	//消费者
	go func() {
		for {
			select {
			case i := <-ch:
				fmt.Printf("%d ", i)
			case <-time.After(3 * time.Second):
				fmt.Println("timeout")
			}
		}
	}()

	time.Sleep(1 * time.Second)
}
