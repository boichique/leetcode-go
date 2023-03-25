package main

import "fmt"

func main() {
	s := "aboba"
	t := "baoba"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	chars := make([]int, 26)

	for i := 0; i < len(s); i++ {
		chars[t[i]-'a']++
		chars[s[i]-'a']--
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}
	}
	return true
}
