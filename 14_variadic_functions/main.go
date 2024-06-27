package main

import "fmt"

func main() {
	VariadicExample("red", "blue", "green", "yellow")
}

func VariadicExample(s ...string) {
	fmt.Println(s[0])
	fmt.Println(s[3])
}
