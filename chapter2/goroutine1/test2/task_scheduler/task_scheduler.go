package task_scheduler

import (
	"fmt"
	"time"
)

func Submit(f func()) {
	go func() {
		defer timer1()()
		f()
	}()
}

func timer1() func() {
	begin := time.Now()
	return func() {
		fmt.Println("执行函数所花时间:", time.Since(begin))
	}
}
