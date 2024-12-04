package main

import (
	"fmt"
)

type Book struct {
	name string
	author string
	year int
}

func (b *Book) getBookInfo() {
	fmt.Printf("Книга с именем: %s\n", b.name)
	fmt.Printf("Книга от автора: %s\n", b.author)
	fmt.Printf("Книга выпущена: %d\n",b.year )
}

func main()  {
	book1:= Book{
		name: "Kniga1",
		author: "Alex",
		year: 1970,
	}

	book2:= Book{
		name: "Kniga2",
		author: "Misha",
		year: 2013,
	}

	book3:= Book{
		name: "Kniga3",
		author: "Pavel",
		year: 2033,
	}

	book1.getBookInfo()
	book2.getBookInfo()
	book3.getBookInfo()
}