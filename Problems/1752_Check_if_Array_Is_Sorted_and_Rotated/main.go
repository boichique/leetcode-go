package main

import "fmt"

func check(nums []int) bool {
	idx := len(nums) - 1
	min := nums[idx]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] <= min && i > 0 && nums[i] == nums[i-1] {
			min = nums[i]
			idx = i - 1
		}
		if nums[i] < min {
			min = nums[i]
			idx = i
		}
	}

	counter := len(nums) - 1
	for counter > 0 {
		if nums[idx%len(nums)] > nums[(idx+1)%len(nums)] {
			return false
		}
		idx++
		counter--
	}
	return true
}

func main() {
	nums := []int{2, 2, 3, 4, 5, 2}
	fmt.Println(check(nums))
}
