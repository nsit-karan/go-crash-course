package main

import "fmt"

func pointerBasics() {
	a := 5
	// address of a
	// Think of this as b pointing to a
	b := &a

	fmt.Println(a, b)

	// will give you int, *int (this is a pointer)
	fmt.Printf("%T, %T\n", a, b)

	// Use * to read value from a address
	// dereference of a pointer
	fmt.Println("Value is", *b)

	// change value using pointer
	// this will change a's value since
	// b is pointing to a !!
	*b = 10
	fmt.Println("Value after change is", a)
}

func main() {
	pointerBasics()
}
