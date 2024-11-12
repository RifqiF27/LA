package handler

import (
	"book-store/collections"
	"book-store/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type BookHandler struct {
	Service service.BookService
	Log     *zap.Logger
}

func NewBookHandler(service service.BookService, log *zap.Logger) *BookHandler {
	return &BookHandler{
		Service: service,
		Log:     log,
	}
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var bookCode, title, fileBook, author, price, discount, cover, categoryID string
	domain := strings.Join([]string{"http://", r.Host}, "")

	bookCode = r.FormValue("book_code")
	title = r.FormValue("title")
	author = r.FormValue("author")
	price = r.FormValue("price")
	discount = r.FormValue("discount")
	categoryID = r.FormValue("category_id")

	categoryIDConvert, _ := strconv.Atoi(categoryID)
	priceConvert, _ := strconv.ParseFloat(price, 64)
	discountConvert, _ :=  strconv.ParseFloat(discount, 64)

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		h.Log.Error("file too large or invalid form data", zap.Error(err))
		http.Error(w, "File too large or invalid form data", http.StatusBadRequest)
		return
	}
	file, data, err := r.FormFile("cover")
	if err != nil {
		h.Log.Error("bookhandler_createBook_formfile", zap.Error(err))
		fmt.Println("Error : ", err)
	}
	defer file.Close()

	dst, err := os.Create(filepath.Join("assets/cover_books", data.Filename))
	if err != nil {
		fmt.Println("Error : ", err)
	}
	_, err = io.Copy(dst, file)
	if err != nil {
		fmt.Println("Error : ", err)
	}

	cover = strings.Join([]string{domain, "/cover_books/", data.Filename}, "")
	fileBook = strings.Join([]string{domain, "/cover_books/", data.Filename}, "")
	
	
	book := collections.Book{
		Category: collections.Category{
			ID: categoryIDConvert, 
		},
		BookCode: bookCode,
		Title:   title,
		Author:  author,
		Price:   float64(priceConvert),
		Discount: float64(discountConvert),
		Cover: cover,
		FileBook: fileBook, 
	}

	if err := h.Service.CreateBook(&book); err != nil {
		h.Log.Error("Failed to create book", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.Service.GetAllBooks()
	if err != nil {
		h.Log.Error("Failed to get books", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Log.Error("Invalid book ID", zap.String("ID", idStr), zap.Error(err))
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := h.Service.GetBookByID(id)
	if err != nil {
		h.Log.Error("Failed to get book by ID", zap.Int("ID", id), zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Log.Error("Invalid book ID", zap.String("ID", idStr), zap.Error(err))
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book collections.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		h.Log.Error("Failed to decode book data", zap.Error(err))
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateBook(id, book); err != nil {
		h.Log.Error("Failed to update book", zap.Int("ID", id), zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book updated successfully"))
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Log.Error("Invalid book ID", zap.String("ID", idStr), zap.Error(err))
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteBook(id); err != nil {
		h.Log.Error("Failed to delete book", zap.Int("ID", id), zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book deleted successfully"))
}
