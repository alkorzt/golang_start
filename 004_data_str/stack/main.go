package main

import (
	"fmt"
	stack "pkg/stacker"
)

func main() {
	var myStack stack.Stack
	myStack.Push("hay")
	myStack.Push(15)
	myStack.Push([]string{"pin", "clip", "needle"})
	for {
		item, err := myStack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}
