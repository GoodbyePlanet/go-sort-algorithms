package utils

func Swap(items []int, i, j int) {
	temp := items[i]
	items[i] = items[j]
	items[j] = temp
}

func GetNumbers() []int {
	return []int{4, 3, 2, 5, 1, 6, 9}
}
