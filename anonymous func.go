package main

import "fmt"

func main() {
	// anonyomous function
	var hi = func() {
		fmt.Println("ha vai kiski")

	}
	var welcome = hi
	// function call
	hi()
	welcome()
}
