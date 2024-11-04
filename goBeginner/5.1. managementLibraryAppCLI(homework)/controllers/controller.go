package controllers

import (
	"errors"
	"fmt"
	"main/models"
	"main/utils"
	"reflect"
	"regexp"
	"strings"
)

func AddBookController(book models.Library) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\033[31mPanic tertangani: Proses tambah buku mengalami kesalahan.\033[0m")
			fmt.Printf("\033[31mError: %v\033[0m\n", r)
		} else {
			fmt.Println("\033[35mProses tambah buku selesai\033[0m")
		}
	}()

	var title, author, isbn string
	var err error

	title, err = utils.InputStr("Masukan title buku:\n")
	if err != nil {
		fmt.Println("\033[31mError:\033[0m", err)
		return
	} else if len(title) < 3 {
		panic(errors.New("\033[31mInput harus lebih dari 3 huruf\033[0m"))
	} else if len(title) > 20 {
		panic(errors.New("\033[31mInput tidak boleh lebih dari 20 huruf\033[0m"))
	}

	author, err = utils.InputStr("Masukan author:\n")
	if err != nil {
		fmt.Println("\033[31mError:\033[0m", err)
		return
	} else if len(author) < 3 {
		panic(errors.New("\033[31mInput harus lebih dari 3 huruf\033[0m"))
	} else if len(author) > 15 {
		panic(errors.New("\033[31mInput tidak boleh lebih dari 15 huruf\033[0m"))
	}

	isbn, err = utils.InputStr("Masukkan no ISBN:\n")
	if err != nil {
		fmt.Println("\033[31mError:\033[0m", err)
		return
	}
	isString := regexp.MustCompile(`^[0-9]+$`).MatchString

	if !isString(isbn) {
		panic(errors.New("\033[31mError: Input harus berupa angka"))
	} else if len(isbn) != 10 && len(isbn) != 13 {
		panic(errors.New("\033[31mError: No ISBN harus 10 angka atau 13 angka"))
	}
	fmt.Println("Data type ISBN : ", reflect.TypeOf(isbn))
	// fmt.Println("Data len : ", len(isbn))
	books := &models.Book{
		Title:  title,
		Author: author,
		ISBN:   isbn,
	}

	if book.IsDuplicate(books) {
		panic(errors.New("\033[31mError: Buku dengan ISBN yang sama sudah ada di perpustakaan\033[0m"))
	}

	book.AddBook(books)
	fmt.Printf("\033[32mBuku berhasil ditambahkan\033[0m\n")

}

func RemoveBookController(book models.Library) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\033[31mPanic tertangani: Proses hapus buku mengalami kesalahan.\033[0m")
			fmt.Printf("\033[31mError: %v\033[0m\n", r)
		} else {
			fmt.Println("\033[35mProses hapus buku selesai\033[0m")
		}
	}()

	var isbn string
	fmt.Print("\033[33mMasukkan No ISBN yang ingin dihapus:\033[0m ")
	fmt.Scanln(&isbn)

	if len(isbn) == 0 {
		panic(errors.New("\033[31mError: ISBN tidak boleh kosong\033[0m"))
	}

	err := book.RemoveBook(isbn)
	if err != nil {
		fmt.Println("\033[31mError:\033[0m", err)
	} else {
		fmt.Printf("\033[32mBuku dengan No ISBN '%s' berhasil dihapus\033[0m\n", isbn)
	}
}

func ShowBookController(book models.Library) {
	fmt.Printf("\033[34mDaftar menu di menu :\033[0m\n\n")
	if len(book.ShowBook()) == 0 {
		fmt.Println("Tidak ada buku di perpustakaan.")
		return
	}
	fmt.Printf("\033[35m%-25s%-20s%-20s\033[0m\n", "Title", "Author", "ISBN")
	fmt.Println(strings.Repeat("-", 60))
	for _, b := range book.ShowBook() {

		fmt.Printf("\033[36m%-25s%-20s%-20s\033[0m\n", b.Title, b.Author, b.ISBN)
		fmt.Println(strings.Repeat("-", 60))

	}
}
