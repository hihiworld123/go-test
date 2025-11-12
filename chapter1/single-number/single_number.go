package main

import "fmt"

func main() {
	nums := []int{2, 2, 1}
	number := singleNumber(nums)
	fmt.Println(number)
}

func singleNumber(nums []int) int {
	data := make(map[int]int)
	for _, v := range nums {
		n, ok := data[v]
		if ok {
			data[v] = n + 1
		} else {
			data[v] = 1
		}
	}

	a := 0
	for k, v := range data {
		if v == 1 {
			a = k
		}
	}

	return a
}
