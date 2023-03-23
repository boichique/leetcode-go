package main

import (
	"fmt"
	"strings"
)

func main() {
	operations := []string{"++X", "++X", "X++"}
	res := finalValueAfterOperations(operations)
	fmt.Println(res)
}

func finalValueAfterOperations(operations []string) int {
	res := 0
	for _, s := range operations {
		if strings.Contains(s, "+") {
			res++
		} else {
			res--
		}
	}
	return res
}
