package models

import "fmt"

type Book struct {
	Title  string
	Author string
	ISBN   string
}

type MyLibrary struct {
	books []Book
}

func (m *MyLibrary) Init() {
	defaultBook := []Book{
		{Title: "Go Programming", Author: "John Doe", ISBN: "1234567890"},
		{Title: "Learning Go", Author: "Tyler A", ISBN: "0987654321"},
		{Title: "Microservices Go", Author: "Diego Rivera", ISBN: "1234567890111"},
		{Title: "Architecture Go", Author: "Eliot Alderson", ISBN: "0987654321000"},
	}

	for _, book := range defaultBook {
		m.AddBook(&book)
	}
}

func (m *MyLibrary) AddBook(b *Book) {
	m.books = append(m.books, *b)
}

func (m *MyLibrary) IsDuplicate(b *Book) bool {
	for _, book := range m.books {
		if book.ISBN == b.ISBN {
			return true
		}
	}
	return false
}

func (m *MyLibrary) RemoveBook(isbn string) error {
	for i, l := range m.books {
		if l.ISBN == isbn {
			m.books = append(m.books[:i], m.books[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("\033[31mBuku dengan no ISBN '%s' tidak ditemukan\033[0m", isbn)

}

func (m *MyLibrary) ShowBook() []Book {
	return m.books
}
