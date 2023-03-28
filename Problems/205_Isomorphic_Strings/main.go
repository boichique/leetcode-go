package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	mapST := map[byte]byte{}
	mapTS := map[byte]byte{}

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		_, check1 := mapST[s[i]]
		_, check2 := mapTS[t[i]]
		if !check1 && !check2 {
			mapST[s[i]] = t[i]
			mapTS[t[i]] = s[i]
		} else if mapST[s[i]] != t[i] || mapTS[t[i]] != s[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("foo", "bar"))
}
