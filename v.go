package main

import "fmt"

func main() {
	var a int = 10
	var b float32 = 10.6
	var name string = "Rajesh"

	fmt.Printf("%v\t %v\t %v\n", a, b, name)

	var i, j string = "Hello", "World"
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", j)

	fmt.Println(i)
	fmt.Println(j)

	fmt.Print(i, "\n", j)
	fmt.Printf("\n%v %v\n", i, j)

	var x, y = 10, "Rajesh"
	fmt.Print(x, y)

	var p = 15.5
	var txt = "Hello !!"

	fmt.Printf("\n%v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%v%%\n", p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)
	fmt.Printf("'%s'\n", txt)
	fmt.Println()

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

}
