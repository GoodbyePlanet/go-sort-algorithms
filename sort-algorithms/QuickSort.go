package main

import (
	"fmt"
	"go-sort-algorithms/utils"
)

func QuickSort(items []int) {

	stack := utils.Stack{}
	stack.Push(0)
	stack.Push(len(items) - 1)

	for !stack.IsEmpty() {

		endIndex := stack.Pop()
		startIndex := stack.Pop()

		partitionElementIndex := partition(items, startIndex, endIndex)

		if partitionElementIndex - 1 > startIndex {
			stack.Push(startIndex)
			stack.Push(partitionElementIndex - 1)
		}

		if partitionElementIndex + 1 < endIndex {
			stack.Push(partitionElementIndex + 1)
			stack.Push(endIndex)
		}
	}

}

func partition(items []int, startIndex, endIndex int) int {

	i := startIndex + 1
	j := endIndex

	for true {

		for items[i] < items[startIndex] {
			i++
			if i == endIndex {
				break
			}
		}

		for items[startIndex] < items[j] {
			j--
			if j == startIndex {
				break
			}
		}

		if i >= j {
			break
		}
		utils.Swap(items, i, j)
	}
	utils.Swap(items, startIndex, j)
	return j
}


func main() {

	numbers := utils.GetNumbers()
	fmt.Println("Unsorted numbers", numbers)

	QuickSort(numbers)
	fmt.Println("Sorted numbers", numbers)
}