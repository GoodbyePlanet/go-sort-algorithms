package utils

func Swap(items []int, i, j int) {
	temp := items[i]
	items[i] = items[j]
	items[j] = temp
}
