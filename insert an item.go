package main
import "fmt"
func main(){
 originalArray:=[]int{1,3,5,6,2}
newitem:=9
index:=3
newArray := insertintoArray(originalArray,newitem,index)
fmt.Println("Original Array:", originalArray)
fmt.Println("Newitem:", newitem)
fmt.Println("Index to insert:", index)
fmt.Println("New Array:", newArrayy)
func insertintoArray(arr[]int,item,index int)[]int{
   newArr:= make([]int,len(arr)+1)
   copy(newArr[index+1:],arr[index:])
   return newArr()

}





}

