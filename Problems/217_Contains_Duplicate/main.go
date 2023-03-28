package main

import "fmt"

func containsDuplicate(nums []int) bool {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
		if freq[num] > 1 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
