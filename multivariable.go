package main

import "fmt"

func main() {

	var a, b = 1, 3.5
	c, d := 5.1, "Ramesh"

	var (
		x int
		y int    = 21
		z string = "Hello, Ramesh!"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println()

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println()

	fmt.Printf("%d\n", x)
	fmt.Printf("%d\n", y)
	fmt.Printf("%s\n", z)
	fmt.Println()

	fmt.Println(x, y, z)
	fmt.Printf("%d %d %s", x, y, z)

}
