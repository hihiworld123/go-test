package main

import (
	"fmt"
	"time"
)

func main() {

	test1()

}

// 编写一个程序，使用 go 关键字启动两个协程，
// 一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行
func test1() {
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				fmt.Printf("奇数：%d\n", i)
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				fmt.Printf("偶数：%d\n", i)
			}
		}
	}()

	time.Sleep(1 * time.Second)
}
