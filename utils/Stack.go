package utils

type Stack struct {
	slice []int
}

func (stack *Stack) Push(item int) {
	stack.slice = append(stack.slice, item)
}

func (stack *Stack) Peek() int {
	return stack.slice[len(stack.slice)-1]
}

func (stack *Stack) Pop() int {
	topItem := stack.Peek()
	stack.slice = stack.slice[0 : len(stack.slice)-1]
	return topItem
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.slice) == 0
}
