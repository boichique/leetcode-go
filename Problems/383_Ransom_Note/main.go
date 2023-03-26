package main

import (
	"fmt"
)

func makeCounter(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range s {
		m[string(v)]++
	}
	return m
}

func canConstruct(ransomNote string, magazine string) bool {
	mapR := makeCounter(ransomNote)
	mapM := makeCounter(magazine)

	for k, v := range mapR {
		if _, exists := mapM[k]; !exists {
			return false
		}
		if mapM[k] < v {
			return false
		}
	}

	return true
}

func main() {
	ransomNote := "aa"
	magazine := "aab"
	fmt.Println(canConstruct(ransomNote, magazine))
}
