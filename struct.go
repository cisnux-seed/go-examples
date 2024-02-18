package main

import "fmt"

func main() {
	var writer Writer
	fmt.Println(writer.Books == nil)
	fmt.Println(writer)

	writer.Name = "Fajra Risqulla"
	writer.EmailAddress = "fajrarisqulla@gmail.com"
	writer.Age = 20
	writer.Books = []Book{
		{
			Title:       "Golang Jr",
			PublishedAt: 2020,
		},
		{
			Title:       "Java Jr",
			PublishedAt: 2010,
		},
	}
	fmt.Println(writer)
	fmt.Println(writer.Name)
	fmt.Println(writer.EmailAddress)
	fmt.Println(writer.Age)
	fmt.Println(writer.Books)
	book1 := Book{
		Title:       "Shingeki no Kyojin",
		PublishedAt: 2012,
	}
	fmt.Println(book1)
	writer.printBooks()
}

type Book struct {
	Title       string
	PublishedAt uint64
}

type Writer struct {
	Name, EmailAddress string
	Age                int
	Books              []Book
}

func (writer Writer) printBooks() {
	for index, book := range writer.Books {
		fmt.Println(index+1, "-", book)
	}
}
