package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i, num := range nums {
		if _, exists := m[num]; exists && i-m[num] <= k {
			return true
		}
		m[num] = i
	}
	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
}
