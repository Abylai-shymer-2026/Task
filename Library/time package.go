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

func (l *Library) Bronn(title string, returnPeriod time.Duration) {
	book := l.Find(title)
	if book == nil {
		fmt.Println("Такой книги нет в библиотеке")
	} else {
		if book.Bron {
			fmt.Printf("Книга уже забронирована до %s\n", book.BronTime.Format("02.01.2006"))
		} else {
			book.Bron = true
			book.BronTime = time.Now().Add(returnPeriod)
			fmt.Printf("Книга успешно забронирована до %s\n", book.BronTime.Format("02.01.2006"))
		}
	}
}

func (l *Library) Info(title string) {
	book := l.Find(title)
	if book == nil {
		fmt.Println("Такой книги нет в библиотеке")
	} else {
		if book.Bron {
			fmt.Printf("Эта книга забронирована до %s\n", book.BronTime.Format("02.01.2006"))
		} else {
			fmt.Println("Эта книга доступна в данный момент")
		}
	}
}

func main() {
	library := Library{}

	library.Add("Абай жолы", "Мухтар Ауезов")
	library.Add("Кытап", "Абылайхан")
	library.Add("Книга", "Абылай")

	library.Bronn("Абай жолы", time.Hour*24*14)
	library.Info("Кытап")
	library.Info("Абай жолы")
	library.Bronn("Книга", time.Hour*24*7)
	library.Info("Книга")
}
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

func (l *Library) Bronn(title string, returnPeriod time.Duration) {
	book := l.Find(title)
	if book == nil {
		fmt.Println("Такой книги нет в библиотеке")
	} else {
		if book.Bron {
			fmt.Printf("Книга уже забронирована до %s\n", book.BronTime.Format("02.01.2006"))
		} else {
			book.Bron = true
			book.BronTime = time.Now().Add(returnPeriod)
			fmt.Printf("Книга успешно забронирована до %s\n", book.BronTime.Format("02.01.2006"))
		}
	}
}

func (l *Library) Info(title string) {
	book := l.Find(title)
	if book == nil {
		fmt.Println("Такой книги нет в библиотеке")
	} else {
		if book.Bron {
			fmt.Printf("Эта книга забронирована до %s\n", book.BronTime.Format("02.01.2006"))
		} else {
			fmt.Println("Эта книга доступна в данный момент")
		}
	}
}

func main() {
	library := Library{}

	library.Add("Абай жолы", "Мухтар Ауезов")
	library.Add("Кытап", "Абылайхан")
	library.Add("Книга", "Абылай")

	library.Bronn("Абай жолы", time.Hour*24*14)
	library.Info("Кытап")
	library.Info("Абай жолы")
	library.Bronn("Книга", time.Hour*24*7)
	library.Info("Книга")
}
