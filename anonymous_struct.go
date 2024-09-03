package main
import "fmt"

type rec func(int,int) int
type rec struct{
	length int
	breadth int
	colour string
	rect rec
}
func main(){
	r := recpara{
		length:5
		breath:20
		colour:"Red",
		rect: func(length int, breadth int)int{
			return length*breadth
		},
	}
	fmt.Println("Acolour of rectangle:" result colour)

	fmt.Println("Area of rectangle is : ", result.rect(result.length, result.breadth))
		}
	





