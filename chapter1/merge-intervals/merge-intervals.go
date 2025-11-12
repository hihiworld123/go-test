package main

import (
	"fmt"
	"sort"
)

func main() {
	//intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	//intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {9, 16}, {15, 18}}
	intervals := [][]int{{4, 7}, {1, 4}}
	data := merge(intervals)
	fmt.Println(data)
}

func merge(intervals [][]int) [][]int {

	//先对数组进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	intervals1 := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		if i == len(intervals)-1 {
			break
		}
		if intervals[i][0] < intervals[i+1][0] &&
			intervals[i][1] >= intervals[i+1][0] &&
			intervals[i][1] < intervals[i+1][1] {
			arr1 := []int{intervals[i][0], intervals[i+1][1]}
			intervals1 = append(intervals1, arr1)
		} else {
			intervals1 = append(intervals1, intervals[i+1])
		}
	}

	return intervals1
}
