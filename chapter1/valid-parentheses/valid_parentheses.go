package main

import (
	"fmt"
	"strings"
)

func main() {
	valid := isValid("([)]")
	fmt.Println(valid)
}

func isValid(s string) bool {

	all := s
	for strings.Contains(all, "()") ||
		strings.Contains(all, "[]") ||
		strings.Contains(all, "{}") {
		all = strings.ReplaceAll(all, "()", "")
		all = strings.ReplaceAll(all, "[]", "")
		all = strings.ReplaceAll(all, "{}", "")
	}

	if len(all) != 0 {
		return false
	}

	return true
}
