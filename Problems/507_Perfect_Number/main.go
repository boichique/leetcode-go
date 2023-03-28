package main

import "fmt"

func checkPerfectNumber(num int) bool {
	sum := 0
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			if i*i != num {
				sum += num / i
			}
		}
	}
	return sum-num == num
}

func main() {
	fmt.Println(checkPerfectNumber(28))
}
