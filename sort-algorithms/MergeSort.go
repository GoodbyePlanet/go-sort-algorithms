package main

import (
	"fmt"
	"go-sort-algorithms/utils"
)

func MergeSort(items []int) {

	for i := 1; i <= len(items)-1; i = i * 2 {
		for j := 0; j < len(items)-1; j += 2 * i {

			mid := j + i - 1
			rightEnd := Min(j+2*i-1, len(items)-1)

			Merge(items, j, mid, rightEnd)
		}
	}
}

func Merge(items []int, leftStart, mid, rightEnd int) {

	aux := make([]int, len(items))
	copy(aux, items)

	i := leftStart
	j := mid + 1

	for k := leftStart; k <= rightEnd; k++ {
		if i > mid {
			items[k] = aux[j]
			j++
		} else if j > rightEnd {
			items[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			items[k] = aux[j]
			j++
		} else {
			items[k] = aux[i]
			i++
		}
	}
}

func Min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func main() {
	numbers := utils.GetNumbers()
	fmt.Println("Unsorted numbers", numbers)

	MergeSort(numbers)
	fmt.Println("Sorted numbers", numbers)
}
