package main

import (
	"fmt"
	"go-sort-algorithms/utils"
)

func bubbleSort(items []int) {

	for i := 0; i < len(items)-1; i++ {
		for j := 0; j < len(items)-1-i; j++ {
			if items[j+1] < items[j] {
				utils.Swap(items, j+1, j)
			}
		}
	}
}

func main() {

	numbers := utils.GetNumbers()
	fmt.Println("Unsorted integers", numbers)

	bubbleSort(numbers)
	fmt.Println("Sorted integers", numbers)
}
