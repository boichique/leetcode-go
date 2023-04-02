package main

import (
	"fmt"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	res := []int{}

	for _, spell := range spells {
		l := 0
		r := len(potions)
		for l != r {
			mid := (l + r) / 2
			if spell*potions[mid] >= int(success) {
				r = mid
			} else {
				l = mid + 1
			}
		}
		res = append(res, len(potions)-r)
	}
	return res
}

func main() {
	spells := []int{5, 1, 3}
	potions := []int{1, 2, 3, 4, 5}
	success := 7
	fmt.Println(successfulPairs(spells, potions, int64(success)))
}
