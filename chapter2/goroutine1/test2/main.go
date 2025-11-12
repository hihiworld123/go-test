package main

import (
	"fmt"
	"go-test/chapter2/goroutine1/test2/task_scheduler"
	"time"
)

// 设计一个任务调度器，接收一组任务（可以用函数表示）
// 并使用协程并发执行这些任务，同时统计每个任务的执行时间。
func main() {

	task_scheduler.Submit(func() {
		fmt.Println("task_scheduler")
		time.Sleep(100 * time.Millisecond)
	})

	time.Sleep(time.Second * 1)
}
