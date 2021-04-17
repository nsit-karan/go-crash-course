package main

import (
	"fmt"
	"strconv"
)

// Define the struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

type PersonConcise struct {
	firstName, lastName string
	age                 int
}

/*
 * Greeting method (value receiver)
 * p here is very similar to this
 * p Person is kind of telling that
 * this function goes with the struct Person
 */
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and i am " + strconv.Itoa(p.age)
}

/*
 * received a pointer
 */
func (p *Person) hasBirthday() {
	p.age++
}

/*
 * getMarried : which will change the last name
 */
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "harry", lastName: "potter", city: "london", gender: "m",
		age: 25}

	// Init person using string : alternate
	person2 := Person{"Ron", "Weasley", "london", "m", 25}

	// Init person using string : alternate
	person3 := Person{"Hermione", "Granger", "london", "f", 25}

	person4 := Person{"Tom", "Riddle", "london", "m", 30}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person1.city)

	// increase age
	person1.age++
	fmt.Println(person1)
	fmt.Println(person1.greet())

	// increase age on a bday
	person1.hasBirthday()
	fmt.Println(person1)

	person3.getMarried("Weasley")
	fmt.Println(person3.greet())

	person4.getMarried("gandhi")
	fmt.Println(person4)

	person5 := PersonConcise{"Dobby", "Elf", 30}
	fmt.Println(person5)
}
