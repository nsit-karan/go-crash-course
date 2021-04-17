package main

import "fmt"

func rangeTutorial() {
	ids := []int{33, 76, 11, 55, 66, 77}
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// without using index, use _ in case not using
	// a paritcular var
	for _, id := range ids {
		fmt.Printf("ID : %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum is", sum)
}

/*
 * range with map
 */
func rangeMap() {
	emails := map[string]string{"harry": "harry@lotr.com", "hermione": "hermione@lotr.com"}
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name:" + k)
	}
}

func main() {
	rangeTutorial()
	rangeMap()
	fmt.Println("done !!")
}
