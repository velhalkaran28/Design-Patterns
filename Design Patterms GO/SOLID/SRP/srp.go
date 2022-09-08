package main

import "fmt"

type Book struct {
	title  string
	author string
}

var books []*Book

func (b *Book) AddBook(bk *Book) {
	books = append(books, bk)
}

func (b *Book) Display() {
	fmt.Println()
	for _, value := range books {
		fmt.Println(fmt.Sprintf("Title = %s, Author = %s", value.title, value.author))
	}
}

// Seperation of Concern
type Upload struct {
	book *Book
}

func UploadToWeb() {
	//
}
func main() {
	b := &Book{}
	b.AddBook(&Book{title: "Design Principles", author: "Karan"})
	b.AddBook(&Book{"Solid Principles", "Velhal"})
	b.Display()
}
