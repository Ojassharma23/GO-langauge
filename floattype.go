/*
float32 float64
complex64(float) complex128
*/
package main

import "fmt"

func main() {
	var student1 string = "John"
	var student2 string = "Jane"
	p := 2

	fmt.Printf("%s\n", student1)
	fmt.Printf("%s\n", student2)
	fmt.Printf("%d\n", p)

	var a float32
	var b float64
	a, b = 255, 254
	fmt.Println("a = %T, b = %T", a, b)
	fmt.Printf("a = %f\n b = %f\n", a, b)
	fmt.Printf("a = %T\n b = %T\n", a, b)
	var x string
	var y int
	var z bool = true

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	var d, e, f = 3, 4, "foo"
	fmt.Println(d, e, f)
	fmt.Println("d is of type %T", d)
	fmt.Println("e is of type %T", e)
	fmt.Println("f is of type %T", f)

}
