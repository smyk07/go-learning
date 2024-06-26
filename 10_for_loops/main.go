package main

import "fmt"

func main() {
	fruits := [3]string{"apple", "orange", "banana"}
	for id, val := range fruits {
		fmt.Printf("%v\t%v\n", id, val)
	}
}
