package main

import "fmt"

func mapTest() {
	emails := make(map[string]string)
	emails["Bob"] = "bob@bob.com"
	emails["harry"] = "harry@harry.com"
	emails["voldemort"] = "voldemort@lotr.com"
	fmt.Println(emails)
	fmt.Printf("Bob's email is : %s\n", emails["Bob"])
	fmt.Printf("len of emails is %d\n", len(emails))

	// delete and element
	delete(emails, "Bob")
	fmt.Println(emails)
}

func mapInitialize() {
	//declare and assing to map
	emails := map[string]string{"Bob": "bob@lotr.com", "harry": "harry@lotr.com"}
	emails["Mike"] = "Mike@lotr.com"
	fmt.Println(emails)
}

func main() {
	mapTest()
	mapInitialize()
	fmt.Println("done")
}
