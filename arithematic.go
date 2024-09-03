// Arithematic operation code in go
package main

// import format libarary
import "fmt"

// constants
const (
	r2     = 1.414
	PI     = 3.14
	C  int = 3     // typed constant (integer)
	D      = "Hi!" //untyped constant (string)
	E      = 20
)

// typed constant
const A int = 1

// untyped constant
const B = 2

// main function
func main() {
	// Initialization
	first_name := "Ojas"
	last_name := "Sharma"
	a := 10
	b := 5
	//  Mainpulations statements
	sum := a + b
	diff := a - b
	prod := a * b
	div := a / b
	// Displaying result
	fmt.Println(first_name, last_name)
	fmt.Println("Addition :", sum)
	fmt.Println("Subtraction :", diff)
	fmt.Println("Multiplication :", prod)
	fmt.Println("Division :", div)
	fmt.Println("Value of Root 2 :", r2)
	area := PI * 5 * 5
	fmt.Println("Area of Circle :", area)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)

}
