/*lvalue and rvalue*/

/*
10							- literal
20							- literal
"Rajesh"	- literal
*/

/*
 a = b -> valid
	b = a -> valid

	b = 10 -> valid
	10 = b -> invalid
	10 = 20 -> invalid
*/

package main

import "fmt"

func main() {

	fmt.Println(15 == 017)    // octal
	fmt.Println(15 == 0xF)    // hexadecimal
	fmt.Println(15 == 0b1111) // binary
	fmt.Println(15 == 15)     // decimal

	var a, b int = 10, 20
	a++ //++a not valid
	b-- //--b not valid
	//pre increment and decrement are not allowed in go
	fmt.Println(a)
	fmt.Println()
	//fmt.Println(b)

	c := 15
	fmt.Printf("Binary of 15 is %0b\n", c)
	fmt.Printf("Octal of 15 is %o\n", c)
	fmt.Printf("Hexadecimal of 15 is %x\n\n", c)

	fmt.Printf("Expression 15==0xF is %T\n", 15 == 0xF)

}
