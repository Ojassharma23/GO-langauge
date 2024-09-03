package main

import "fmt"

type book struct {
	tittle  string
	book    string
	author  string
	book_id string
}

func main() {
	var book1 book
	book1.tittle = "go lang book"
	book1.author = "mahesh babu"
	book1.book = "programming  in go language"
	book1.book_id = "007"

	//accessing the fields of a structure using dot operator
	fmt.Println("Title : ", book1.tittle)
	fmt.Println("Book Type: ", book1.book)
	fmt.Println("Author Name : ", book1.author)
	fmt.Println("Book_id : ", book1.book_id)

}
