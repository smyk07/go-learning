package main

import "fmt"

func main() {
	x := 17

	if x > 10 {
		fmt.Println("x is greater than 10")
		if x > 15 {
			fmt.Println("x is greater than 15")
		} else {
			fmt.Println("x is smaller than 15")
		}
	} else {
		fmt.Println("x is smaller than 10")
	}
}
