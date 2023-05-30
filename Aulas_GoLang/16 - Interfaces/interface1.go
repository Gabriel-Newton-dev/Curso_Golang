package main

import "fmt"

type Book struct {
	Title  string
	Author string
}

type Stringer interface {
	Article() string
}

func (b Book) Article() string {
	return fmt.Sprintf("This book %s was written by %s", b.Title, b.Author)
}

func main() {

	books := Book{"The Go programming Language", "Brian W. Kernighan"}
	fmt.Println(books.Article())

}
