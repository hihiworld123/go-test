package main

import (
	"fmt"
	"strconv"
)

func main() {
	palindrome := isPalindrome(121)
	fmt.Println(palindrome)
}
func isPalindrome(x int) bool {
	str1 := strconv.Itoa(x)
	arr := []rune(str1)
	len1 := len(arr)
	len2 := (len1 / 2) + 1
	for i := 0; i < len2; i++ {
		if arr[i] != arr[len1-i-1] {
			return false
		}
	}

	return true
}
