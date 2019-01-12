package main

import (
	"fmt"
	"go-sort-algorithms/utils"
)

func insertionSort(items []int) {

	for i := 0; i < len(items); i ++ {
		for j := i; j > 0; j-- {
			if items[j] < items[j-1] {
				utils.Swap(items, j, j-1)
			} else {
				break
			}
		}
	}
}

func main() {

	numbers := utils.GetNumbers()
	fmt.Println("Unsorted integers", numbers)

	insertionSort(numbers)
	fmt.Println("Sorted numbers", numbers)
}
