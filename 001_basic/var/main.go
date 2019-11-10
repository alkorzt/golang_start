package main

import (
	"fmt"
	"unicode/utf8"
)

// UserID ...
type UserID int

const (
	x  = iota
	x1 = 5 * iota
	x2
)

func main() {
	// var intVar0 int
	// var intVar1 int = 1
	// var intVarDyn = 42
	// varShort := 24
	// var1, var2 := 1, 2
	// var s string = "test"
	// var ch byte = 'a'
	var strHello = "Привет"

	fmt.Println(len(strHello))
	fmt.Println(utf8.RuneCountInString(strHello))

	fmt.Println(x, x1, x2)

	mID := UserID(1)
	println(mID)
}
