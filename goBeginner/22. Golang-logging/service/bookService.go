package service

import (
	"book-store/collections"
	"book-store/repository"
	"fmt"

	"go.uber.org/zap"
)

type BookService interface {
	CreateBook(book *collections.Book) error
	GetAllBooks() ([]collections.Book, error)
	GetBookByID(id int) (*collections.Book, error)
	UpdateBook(id int, book collections.Book) error
	DeleteBook(id int) error
}

type BookServiceImpl struct {
	Repo repository.BookRepository
	Log  *zap.Logger
}

func NewBookService(repo repository.BookRepository, log *zap.Logger) BookService {
	return &BookServiceImpl{
		Repo: repo,
		Log:  log,
	}
}

func (s *BookServiceImpl) CreateBook(book *collections.Book) error {
	s.Log.Info("Creating new book", zap.String("BookCode", book.BookCode))
	err := s.Repo.Create(book)
	if err != nil {
		s.Log.Error("Failed to create book", zap.Error(err))
		return fmt.Errorf("could not create book: %w", err)
	}
	return nil
}

func (s *BookServiceImpl) GetAllBooks() ([]collections.Book, error) {
	s.Log.Info("Fetching all books")
	books, err := s.Repo.GetAll()
	if err != nil {
		s.Log.Error("Failed to fetch books", zap.Error(err))
		return nil, fmt.Errorf("could not fetch books: %w", err)
	}
	return books, nil
}

func (s *BookServiceImpl) GetBookByID(id int) (*collections.Book, error) {
	s.Log.Info("Fetching book by ID", zap.Int("ID", id))
	book, err := s.Repo.GetByID(id)
	if err != nil {
		s.Log.Error("Failed to fetch book by ID", zap.Error(err))
		return nil, fmt.Errorf("could not fetch book by ID: %w", err)
	}
	return book, nil
}

func (s *BookServiceImpl) UpdateBook(id int, book collections.Book) error {
	s.Log.Info("Updating book", zap.Int("ID", id))
	err := s.Repo.Update(id, book)
	if err != nil {
		s.Log.Error("Failed to update book", zap.Error(err))
		return fmt.Errorf("could not update book: %w", err)
	}
	return nil
}

func (s *BookServiceImpl) DeleteBook(id int) error {
	s.Log.Info("Deleting book", zap.Int("ID", id))
	err := s.Repo.Delete(id)
	if err != nil {
		s.Log.Error("Failed to delete book", zap.Error(err))
		return fmt.Errorf("could not delete book: %w", err)
	}
	return nil
}
