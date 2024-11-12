package repository

import (
	"book-store/collections"
	"database/sql"
	"fmt"

	"go.uber.org/zap"
)

type BookRepository interface {
	Create(book *collections.Book) error
	GetAll() ([]collections.Book, error)
	GetByID(id int) (*collections.Book, error)
	Update(id int, book collections.Book) error
	Delete(id int) error
}

type BookRepositoryDb struct {
	DB *sql.DB
	Log *zap.Logger
}

func NewBookRepository(db *sql.DB, log *zap.Logger) BookRepository {
	return &BookRepositoryDb{
		DB: db,
		Log: log,
	}
}

func (b *BookRepositoryDb) Create(book *collections.Book) error {
	query := `INSERT INTO "Books" (book_code, title, category_id, author, price, discount, cover, file_book) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := b.DB.QueryRow(query, book.BookCode, book.Title, book.Category.ID, book.Author, book.Price, book.Discount, book.Cover, book.FileBook).Scan(&book.ID)
	if err != nil {
		b.Log.Error("failed to create book", zap.Error(err))
		return fmt.Errorf("failed to create book: %w", err)
	}
	return nil
}

func (b *BookRepositoryDb) GetAll() ([]collections.Book, error) {
	var books []collections.Book
	query := `SELECT id, book_code, title, category_id, author, price, discount, cover, file_book FROM "Books"`
	rows, err := b.DB.Query(query)
	if err != nil {
		b.Log.Error("failed to fetch books", zap.Error(err))
		return nil, fmt.Errorf("failed to fetch books: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var book collections.Book
		if err := rows.Scan(&book.ID, &book.BookCode, &book.Title, &book.Category.ID, &book.Author, &book.Price, &book.Discount, &book.Cover, &book.FileBook); err != nil {
			b.Log.Error("failed to scan book", zap.Error(err))
			return nil, fmt.Errorf("failed to scan book: %w", err)
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		b.Log.Error("error iterating books rows", zap.Error(err))
		return nil, fmt.Errorf("error iterating books rows: %w", err)
	}

	return books, nil
}

func (b *BookRepositoryDb) GetByID(id int) (*collections.Book, error) {
	var book collections.Book
	query := `SELECT id, book_code, title, category_id, author, price, discount, cover, file_book FROM "Books" WHERE id=$1`
	err := b.DB.QueryRow(query, id).Scan(&book.ID, &book.BookCode, &book.Title, &book.Category.ID, &book.Author, &book.Price, &book.Discount, &book.Cover, &book.FileBook)
	if err != nil {
		b.Log.Error("failed to get book by ID", zap.Error(err))
		return nil, fmt.Errorf("failed to get book by ID: %w", err)
	}
	return &book, nil
}

func (b *BookRepositoryDb) Update(id int, book collections.Book) error {
	query := `UPDATE "Books" SET book_code=$1, title=$2, category_id=$3, author=$4, price=$5, discount=$6, cover=$7, file_book=$8 WHERE id=$8`
	_, err := b.DB.Exec(query, book.BookCode, book.Title, book.Category.ID, book.Author, book.Price, book.Discount, book.Cover, book.FileBook, id)
	if err != nil {
		b.Log.Error("failed to update book", zap.Error(err))
		return fmt.Errorf("failed to update book: %w", err)
	}
	return nil
}


func (b *BookRepositoryDb) Delete(id int) error {
	query := `DELETE FROM "Books" WHERE id=$1`
	_, err := b.DB.Exec(query, id)
	if err != nil {
		b.Log.Error("failed to delete book", zap.Error(err))
		return fmt.Errorf("failed to delete book: %w", err)
	}
	return nil
}
