package main

import "fmt"

// 指针练习
func main() {
	a := 1
	test1(&a)
	fmt.Println(a)

	arr := []int{1, 3, 5}
	test2(arr)
	fmt.Println(arr)
}

// 题目 ：编写一个Go程序，定义一个函数，
// 该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，
// 然后在主函数中调用该函数并输出修改后的值。
func test1(a *int) {
	*a += 10
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
func test2(arr []int) {
	for i := range arr {
		arr[i] = arr[i] * 2
	}
}
