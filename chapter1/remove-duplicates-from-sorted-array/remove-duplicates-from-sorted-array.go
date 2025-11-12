package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	duplicates := removeDuplicates(nums)
	fmt.Println(duplicates)
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[index] && i < len(nums)-1 {
			for j := i; j < len(nums)-1; j++ {
				nums[j] = nums[j+1]
			}
			index++
		}
	}

	total := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[len(nums)-1] {
			total++
		}
	}

	return total
}
