/*
uint8 uint16 uint32 uint64
int8 int16 int32 int64
*/
package main

import "fmt"

func main() {
	var a uint8
	var b uint8
	a, b = 255, 254
	fmt.Printf("a = %d\n b = %d\n", a, b)
}
