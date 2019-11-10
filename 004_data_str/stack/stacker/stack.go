package stack

import "errors"

// Stack  - тип стека
type Stack []interface{}

// Len - длина стека
func (stack Stack) Len() int {
	return len(stack)
}

// Push - вставить элемент в стек
func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}

// Top - элемент с вершины стека
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("stack empty")
	}
	return stack[len(stack)-1], nil
}

// Pop - извлечь элемент из стека
func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("empty stack")
	}
	x := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return x, nil
}
