package main

import (
	"fmt"
)

func main() {
	x := 1234321
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reversedX := 0
	temp := x
	for temp > 0 {
		reversedX = reversedX*10 + temp%10
		temp /= 10
	}
	return x == reversedX
}
