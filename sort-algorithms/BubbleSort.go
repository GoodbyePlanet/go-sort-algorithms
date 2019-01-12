package main

import (
	"fmt"
	"go-sort-algorithms/utils"
)

func sort(items []int) {

	for i := 0; i < len(items)-1; i++ {
		for j := 0; j < len(items)-1-i; j++ {
			if items[j+1] < items[j] {
				utils.Swap(items, j+1, j)
			}
		}
	}
}

func main() {

	numbers := []int{4, 3, 2, 5, 1, 6, 9}
	fmt.Println("Unsorted integers", numbers)

	sort(numbers)
	fmt.Println("Sorted integers", numbers)
}
