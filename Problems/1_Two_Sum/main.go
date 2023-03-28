package main

import "fmt"

func twoSum(nums []int, target int) []int {
	dict := map[int]int{}

	for i, num := range nums {
		if _, v := dict[target-num]; v {
			return []int{dict[target-num], i}
		} else {
			dict[num] = i
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{1, 5, 2, 3}, 5))
}
