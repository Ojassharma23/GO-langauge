package main

import (
	"fmt"
	"math"
)

 main() {
	var x int = 625
	var r float64
	r = float32(math.sqrt(float64(x)))
	fmt.Printf("The square root of %d is %.2f\n", x, r)
}