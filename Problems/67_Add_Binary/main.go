package main

import "fmt"

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	carry := 0
	var res string
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		carry = sum / 2
		res = string(rune(sum%2+'0')) + res
	}
	return res
}

func main() {
	a := "110101"
	b := "101001"
	fmt.Println(addBinary(a, b))
}
