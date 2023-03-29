package main

import "fmt"

func flipAndInvertImage(image [][]int) [][]int {
	for i := 0; i < len(image); i++ {
		for l, r := 0, len(image[i])-1; l <= r; l, r = l+1, r-1 {
			image[i][l], image[i][r] = (image[i][r]+1)%2, (image[i][l]+1)%2
		}
	}
	return image
}

func main() {
	fmt.Println(flipAndInvertImage([][]int{{0, 1, 1}, {0, 0, 1}, {0, 1, 0}}))
}
