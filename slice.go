// package main

// import "fmt"
// func main(){
// 	var numbers  = make([]int,4,5)
// 	printSlice(numbers)
// }
// func printSlice
// fmt.Println("n",numbers)

//sorting the slice
// package main
// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	slice := []int{7, 2, 3, 9, 1}
// 	sort.Ints(slice)
// 	fmt.Println("sorted slice :")
// 	for _, item := range slice {
// 		fmt.Printf("%d ", item)
// 	}
// }

// the slice is sort or not
package main
import (
	"fmt"
	"sort"
)

func main() {
	var status bool = false
	slice1 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	slice2 := []int{20, 30, 58, 11, 99, 20, 5, 79}

	status = sort.IntsAreSorted(slice1)
	if status == true {
		fmt.Println("The first Slice is sorted.")
	} else {
		fmt.Println("The first Slice is not sorted.")
	}
	status = sort.IntsAreSorted(slice2)
	if status == true {
		fmt.Println("\nThe second Slice is sorted.\n")
	} else {
		fmt.Println("\nThe second Slice is not sorted.\n")
	}
}
