package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := []string{"flower", "flow", "flight"}
	prefix := longestCommonPrefix(arr)
	fmt.Println(prefix)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	str0 := strs[0]
	length := len(str0)
	prefix := ""

	strLength := len(strs)
	for i := 1; i < length; i++ {
		str := str0[:i]
		count := 0
		for _, str1 := range strs {
			if strings.HasPrefix(str1, str) {
				count++
			}
		}
		if count != strLength {
			break
		}
		prefix = str
	}

	return prefix
}
