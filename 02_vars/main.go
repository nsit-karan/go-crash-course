package main

import "fmt"

func main() {
	var name string = "Harry"
	var age1 int = 20
	fmt.Println(name, age1)

	var name2 = "Ron"
	var age2 int32 = 21
	var isCool bool = true
	fmt.Println(name2, age2)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", isCool)

	const name3 = "Hermione"
	// Not allowed : changing constant name3 = "Malfoy"

	// short way to assign values to var
	name4 := "Lupin"
	fmt.Println(name4)

	var size float32 = 4.5
	fmt.Printf("%T\n", size)

	// Assign stuff together
	name5, email5 := "snape", "snap5@hogwarts.com"
	fmt.Println(name5, email5)
}
