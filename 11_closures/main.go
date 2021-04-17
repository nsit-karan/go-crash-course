package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// sum_func holds the function itself.
	// think of this like sum_func being an alias to adder
	// the variable sum inside adder keeps on holding the value
	// across calls to adder
	// since the value is help by the sum_func
	sum_func := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(sum_func(i))
	}

	fmt.Println("done")
}
