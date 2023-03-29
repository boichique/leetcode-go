package main

import "fmt"

func containsDuplicate(nums []int) bool {
    freq := make(map[int]bool)
    for _, num := range nums {
        if freq[num] {
            return true
        }
        freq[num] = true
    }
    return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
