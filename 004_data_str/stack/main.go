package main

import (
	"fmt"
	stack "go_start/004_data_str/stack/stacker"
)

func main() {
	var myStack stack.Stack
	myStack.Push("hay")
	myStack.Push(15)
	myStack.Push([]string{"pin", "clip", "needle"})
	fmt.Println(myStack.Top())
	for {
		item, err := myStack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}
