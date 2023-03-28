package main

import (
	"fmt"
	"strings"
)

func getValuesFormMap(m map[byte]string) string {
	var v string
	for _, value := range m {
		v += value
	}
	return v
}

func wordPattern(pattern string, s string) bool {
	m := map[byte]string{}
	splitedS := strings.Split(s, " ")

	if len(splitedS) != len(pattern) {
		return false
	}

	for i := 0; i < len(splitedS); i++ {
		if _, exist := m[pattern[i]]; exist && m[pattern[i]] != splitedS[i] {
			return false
		} else if !exist && strings.Contains(getValuesFormMap(m), splitedS[i]) {
			return false
		}
		m[pattern[i]] = splitedS[i]
	}
	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
}
