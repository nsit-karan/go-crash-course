package main

import "fmt"

func numComparison(x int, y int) {
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d", y, x)
	}
}

func colorCheck(color string) {
	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Printf("color is not blut or red, its %s", color)
	}
}

func colorCheckSwitch(color string) {
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("Color is neither blue or red")
	}
}

func main() {
	x := 5
	y := 10
	color := "red"

	numComparison(x, y)
	colorCheck(color)
	colorCheckSwitch(color)
}
