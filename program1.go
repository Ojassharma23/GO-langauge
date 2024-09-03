package main

import (
	"fmt"
)

func main() {
	name := "Prathamesh Ratthe"
	age := 20
	branch := "Cyber Security"
	college := "Shri Ramdeobaba College of Engineering & Management"
	fmt.Printf("Name: %s \t Age: %d\n", name, age)
	fmt.Printf("Branch: %s \t\t College: %s\n", branch, college)

	var a string = "Rajesh"
	fmt.Printf("Value of a is '%s'\n", a)
	fmt.Println("Value of a is " + "'" + a + "'")

	const n, dept = "GeeksforGeeks", "CS"
	fmt.Printf("%s is a %s Portal.\n", n, dept)

	var str = "GeeksforGeeks"
	fmt.Printf("The string is %s \n", str)

	var num1 int = 21
	fmt.Printf("The decimal value is %d \n", num1)

	var num2 float32 = 7.786
	fmt.Printf("The floating point is %g \n", num2)

	var num3 int = 14
	fmt.Printf("The binary value of num3 is %b \n", num3)

	var num4 = 4 + 1i
	fmt.Printf("Scientific notation of num4 : %e \n", num4)

}
