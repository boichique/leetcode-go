package main

import "fmt"

func sumOfUnique(nums []int) int {
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}
	count := 0
	for k, v := range m {
		if v == 1 {
			count += k
		}
	}
	return count
}

func main() {
	fmt.Println(sumOfUnique([]int{1, 2, 3, 4, 5, 5}))
}
