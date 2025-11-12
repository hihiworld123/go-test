package main

import "fmt"

func main() {
	nums := []int{3, 3}
	target := 6
	indexArr := twoSum(nums, target)
	fmt.Println(indexArr)
}

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
