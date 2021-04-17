package main

import "fmt"

func main() {
	//Arrays
	var fruitArray [2]string

	//Assign values
	fruitArray[0] = "apple"
	fruitArray[1] = "orange"

	// Declare and assign
	vegArray := [2]string{"lauki", "cauliflower"}

	fmt.Println(fruitArray)
	fmt.Println(vegArray)

	// array slice
	fruitSlice := []string{"grape", "watermelon", "banana"}
	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[1:2])
}
