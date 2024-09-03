// package main

// import "fmt"

// func main() {
// 	// create a map
// 	submarks := map[string]int{"golang": 100, "java": 80, "python": 95}
// 	fmt.Println(submarks)
// }
///////////////////////////////////////////////////////////////////////////////////////////////////

// package main

// import "fmt"

// func main() {
// 	fc := map[string]string{"sunflower": "yellow", "rose": "red", "lotus": "pink"}
// 	// access the value for the sunflower
// 	fmt.Println(fc["sunflower"])
// 	fmt.Println(fc["lotus"])
// }

///////////////////////////////////////////////////////////////////////////////////////////////////

// package main

// import "fmt"

// func main() {
// 	capital := map[string]string{"india": "mumbai"}
// 	fmt.Println("intial map:", capital)
// 	capital["india"] = "delhi"
// 	fmt.Println("final map:", capital)
// }

///////////////////////////////////////////////////////////////////////////////////////////////////
// adding the value in map
// package main

// import "fmt"

// func main() {
// 	S := map[int]string{1: "ojas", 2: "sharma"}
// 	fmt.Println("map", S)
// 	// add the element in the 3
// 	S[3] = "vishnu"
// 	// add the element in the 5
// 	S[5] = "kamal"
// 	fmt.Println("new map", S)
// }

// /////////////////////////////////////////////////////////////////////////////////////////////////
// delete the value in slice
// package main

// import "fmt"

// func main() {
// 	S := map[int]string{1: "ojas", 2: "sharma"}
// 	fmt.Println("", S)
// // deleting the value in slice
// 	delete(S, 2)

// 	fmt.Println("new", S)
// }

package main

import "fmt"

func main() {
	squarednumber := map[int]int{2: 4, 3: 9, 4: 16}
	// for range loop to iterrate through each key vakue in map
	for number.squared := range squarednumber {
		fmt.Printf("Square of the %d is %d\n ", number, squared)
	}
}
