package models

type Library interface {
	AddBook(Book *Book)
	RemoveBook(ISBN string) error
	ShowBook() []Book
	IsDuplicate(b *Book) bool
}
