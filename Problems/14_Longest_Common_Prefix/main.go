package main

import (
	"fmt"
)

func findStringWitnShortestLen(s []string) string {
	if len(s) < 2 {
		return s[0]
	}
	min := len(s[0])
	var shortest string
	for _, v := range s {
		if len(v) <= min {
			min = len(v)
			shortest = v
		}
	}
	return shortest
}

func longestCommonPrefix(strs []string) string {
	shortest := findStringWitnShortestLen(strs)
	for i, ch := range shortest {
		for _, s := range strs {
			if s[i] != byte(ch) {
				return shortest[:i]
			}
		}
	}
	return shortest
}

func main() {
	s := []string{"cir", "car"}
	fmt.Println(longestCommonPrefix(s))
}
