package main

import "fmt"

type Person struct {
	name string
	age  int
}

func print_person(pers Person) {
	fmt.Println("Name:", pers.name)
	fmt.Println("Age:", pers.age)
}

func main() {
	var pers1 Person
	pers1.name = "Clix"
	pers1.age = 112358

	print_person(pers1)
}
