package main

import "fmt"

func strStr(haystack string, needle string) int {
	l := 0
	r := len(needle)

	for r <= len(haystack) {
		if haystack[l:r] == needle {
			return l
		} else {
			l++
			r++
		}
	}
	return -1
}

func main() {
	haystack := "sabbuttrue"
	needle := "but"
	fmt.Println(strStr(haystack, needle))
}
