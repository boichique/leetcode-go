package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	res := s
	for strings.Contains(res, "[]") || strings.Contains(res, "()") || strings.Contains(res, "{}") {
		replacer := strings.NewReplacer("{}", "", "()", "", "[]", "")
		res = replacer.Replace(res)
	}
	return len(res) == 0
}

func main() {
	fmt.Println(isValid("({[([{}])]})()"))
}
