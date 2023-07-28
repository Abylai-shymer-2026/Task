package main

import (
	"fmt"
	"time"
)

type Book struct {
	Title    string
	Author   string
	Bron     bool
	BronTime time.Time
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
func (l *Library) Info(title string) {
	book := l.Find(title)
	if book == nil {
		fmt.Println("Такой книги нет в библиотеке")
	} else {
		if book.Bron {
			fmt.Printf("Книга сейчас у %s до %s\n", book.Author, book.BronTime)
		} else {
			fmt.Println("Это книга данный момент свободен")
		}
	}
}

func main() {

	library := Library{}

	library.Add("Абай жолы", "Мухтар Ауезов")
	library.Add("Книга", "Абылайхан")
	library.Add("Кытап", "Абылай")

	library.Bronn("Абай жолы")
	library.Bronn("Кытап")
	library.Bronn("Кытап")

	library.Info("Абай жолы")
	library.Info("Кытап")
	library.Info("Книга")
}
