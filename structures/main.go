package main

import (
	"fmt"
	"structures/book"
)

func main() {
	book1 := book.FantasyBook{"J.K. Rowling", "Harry Potter", 56}
	greeting := book.SayHi()
	fmt.Println(greeting)
	fmt.Println(book1.CopiesSold)
	book1.CopiesSold = 50
	fmt.Println(book1.CopiesSold)
}
