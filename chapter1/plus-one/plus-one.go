package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	arr := []int{9}
	one := plusOne(arr)
	fmt.Println(one)
}

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}

	str := ""
	for _, v := range digits {
		str += strconv.Itoa(v)
	}

	atoi, _ := strconv.Atoi(str)
	atoi = atoi + 1

	itoa := strconv.Itoa(atoi)
	split := strings.Split(itoa, "")

	arr := []int{}
	for _, s := range split {
		num, _ := strconv.Atoi(s)
		arr = append(arr, num)
	}

	return arr
}
