package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Bron   bool
}

type Library struct {
	Bookbook []Book
}

func (l *Library) Add(title, author string) {
	book := Book{Title: title, Author: author, Bron: false}
	l.Bookbook = append(l.Bookbook, book)
}

func (l *Library) Find(title string) *Book {
	for i := range l.Bookbook {
		if l.Bookbook[i].Title == title {
			return &l.Bookbook[i]
		}
	}
	return nil
}

func (l *Library) Bronn(title string) {
	book := l.Find(title)
	if book == nil {
		fmt.Println("Такой книги нет в библиотеке")
	} else {
		if book.Bron {
			fmt.Println("Книга уже забронирована")
		} else {
			book.Bron = true
			fmt.Println("Книга успешно забронирована")
		}
	}
}

func main() {

	library := Library{}

	library.Add("Абай жолы", "Мухтар Ауезов")
	library.Add("Книга", "Абылайхан")
	library.Add("Кытап", "Абылай")

	library.Bronn("Абай жолы")
	library.Bronn("Book")
	library.Bronn("Кытап")
}
