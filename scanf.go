package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello", name)

	var fname, lname string
	fmt.Println("Enter your first name and last name: ")
	fmt.Scanln(&fname, &lname)
	fmt.Println("Hello", fname, lname)

	fmt.Print("Enter fisrt string: ")
	var first string
	fmt.Scanln(&first)

	fmt.Print("Enter second string: ")
	var second string
	fmt.Scanln(&second)

	fmt.Print(first + " " + second)

	var a, b int = 10, 20
	fmt.Println(a, b)
	fmt.Println(a + b)

}
