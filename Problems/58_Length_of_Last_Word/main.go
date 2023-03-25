package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	splited := strings.Split(strings.Trim(s, " "), " ")
	return len(splited[len(splited)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("My name is Ivan"))
}
