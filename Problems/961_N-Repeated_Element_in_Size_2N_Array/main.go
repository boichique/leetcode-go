package main

import "fmt"

func repeatedNTimes(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] == len(nums)/2 {
			return v
		}
	}
	return -1
}

func main() {
	nums := []int{2, 5, 7, 7}
	fmt.Println(repeatedNTimes(nums))
}
