package main

import "fmt"

func main() {

	var k = 15
	fmt.Printf("%b\n", k)
	fmt.Printf("%d\n", k)
	fmt.Printf("%+d\n", k) // + se sign print krega
	fmt.Printf("%o\n", k)
	fmt.Printf("%O\n", k)
	fmt.Printf("%x\n", k)
	fmt.Printf("%X\n", k)
	fmt.Printf("%#x\n", k) // # will print in 0x format
	fmt.Printf("%4d\n", k)
	fmt.Printf("%-4d\n", k)
	fmt.Printf("%05d\n", k)

	var txt = "Hello"
	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

}
