package main

import (
	"fmt"
	"go-sort-algorithms/utils"
)

func SelectionSort(items []int) {

	for i := 0; i < len(items); i++ {
		currentMinIndex := GetCurrentMinIndex(items, i)
		if items[currentMinIndex] < items[i] {
			utils.Swap(items, i, currentMinIndex)
		}
	}
}

func GetCurrentMinIndex(items []int, index int) int {
	min := index

	for i := index + 1; i < len(items); i++ {
		if items[i] < items[min] {
			min = i
		}
	}
	return min
}

func main() {

	numbers := utils.GetNumbers()
	fmt.Println("Unsorted integers", numbers)

	SelectionSort(numbers)
	fmt.Println("Sorted integers", numbers)
}
